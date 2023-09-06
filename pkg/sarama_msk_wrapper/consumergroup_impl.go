package sarama_msk_wrapper

import (
	"context"
	"time"

	"github.com/rs/zerolog"
)

func (c1 *consumerWrapperImpl) Start() {
	go c1.startSync()
}

func (c1 *consumerWrapperImpl) startSync() {

	logger := getLoggerWithName(_packageNameConsumerGroup + ":startSync()")
	// Temporary set the level to INFO when starting
	logger = logger.Level(zerolog.InfoLevel)
	fields := map[string]interface{}{"config": c1.config.string()}

	if c1.hasStarted.Load() {
		logger.Info().Fields(fields).Msg("ConsumerGroupId has already been started!")
		return
	}
	c1.hasStarted.Store(true)
	logger.Info().Fields(fields).Msg("Starting ConsumerGroupId.")
	// Reset logger level now
	logger = logger.Level(_logLevel)

	ctx := context.Background()
	newCtx, cancel := context.WithCancel(ctx)
	// We need to close the consumer group in a defer call
	defer func() {
		logger.Info().Fields(fields).Msg("Cancelling context")
		cancel()

		logger.Info().Fields(fields).Msg("Closing consumer group")
		_ = c1.consumerGroup.Close()
	}()

	// Track errors
	go func() {
		ticker := time.NewTicker(c1.durationToResetCounter)

		defer func() {
			logger.Info().Fields(fields).Msg("Stopping ConsumerGroupId ticker.")
			ticker.Stop()
		}()

		// TODO: Write check for not consuming within x minutes / hours / etc.
	ErrorLoop:
		for {
			select {
			case consumerError := <-c1.consumerGroup.Errors():
				logger.Error().Fields(fields).Err(consumerError).Msg("error consuming message")
				c1.errorCount.Add(1)
				c1.funcMetricErrorConsuming()
				c1.funcErrorHandling(consumerError)
			case <-ticker.C:
				errorCountStr := _printer.Sprintf(_formatDigit, c1.errorCount.Load())
				logger.Info().Fields(fields).Stringer("durationToReset", c1.durationToResetCounter).
					Str("errorCount", errorCountStr).Msg("resetting counter for ConsumerGroup")
				c1.resetCount()
			case <-c1.stopChan:
				break ErrorLoop
			}
		}
	}()

	// MAIN: Start the consumer!
	// Iterate over consumer sessions in an infinite loop, as suggested by Sarama.
	// Ref: https://pkg.go.dev/github.com/Shopify/sarama#ConsumerGroup
	// Ref: https://github.com/Shopify/sarama/blob/main/examples/consumergroup/main.go#L102
	go func() {
	ConsumeLoop:
		for {
			err := c1.consumerGroup.Consume(newCtx, c1.topics, c1.consumerGroupHandlerWrapper)
			if err != nil {
				logger.Error().Fields(fields).Err(err).
					Msg("error encountered when ConsumerGroup tried to consume")
				break ConsumeLoop
			}
			// Check if context was cancelled, signaling that the consumer should stop
			if newCtx.Err() != nil {
				break ConsumeLoop
			}
		}
	}()

	// Await until stopped
	<-c1.stopChan
}

func (c1 *consumerWrapperImpl) Stop() {
	logger := getLoggerWithName(_packageNameConsumerGroup + ":Stop()")
	fields := map[string]interface{}{
		"config":            c1.config.string(),
		"_terminationDelay": _terminationDelay.String(),
	}
	logger.Info().Fields(fields).Msg("Stopping ConsumerWrapper using `stopChan`")

	if c1.hasStopped.Load() {
		logger.Info().Fields(fields).Msg("WARNING! ConsumerWrapper is already stopped.")
		return
	}

	c1.hasStopped.Store(true)
	// Send stop signal to channel. This must be the last call in the function, otherwise the send might prevent
	// other statements from executing.
	close(c1.stopChan)

	_ = _wr.Close()

	time.Sleep(_terminationDelay)
	deleteTmpCerts()
	logger.Info().Fields(fields).Msg("Finished termination delay")
}

func (c1 *consumerWrapperImpl) resetCount() {
	c1.errorCount.Store(0)
}
