package sarama_msk_wrapper

import (
	"errors"
	"strconv"
	"time"

	"github.com/IBM/sarama"
	"github.com/rs/zerolog"
	"github.com/toddnguyen47/util-go/pkg/timeisoparser"
)

func (a1 *asyncProducerImpl) PublishMessage(message sarama.ProducerMessage) error {

	logger := getLoggerWithName(_packageNameAsyncProducer + ":PublishMessage()")
	b1, err := message.Key.Encode()
	if err != nil {
		b1 = []byte(strconv.Itoa(message.Key.Length()))
	}
	fields := map[string]interface{}{
		"topic":      message.Topic,
		"messageKey": string(b1),
		"value_len":  message.Value.Length(),
		"config":     a1.config.string(),
	}

	if !a1.hasStarted.Load() {
		msg := "ERROR! AsyncProducer has not been started"
		logger.Error().Fields(fields).Msg(msg)
		return errors.New(msg)
	}

	logger.Info().Fields(fields).Msg("INIT PublishMessage()")

	if a1.HasClosed() {
		msg := "ERROR! AsyncProducer already been closed"
		logger.Error().Fields(fields).Msg(msg)
		return errors.New(msg)
	}

	a1.asyncProducer.Input() <- &message
	a1.enqueuedCount.Add(1)
	logger.Info().Fields(fields).Msg("END PublishMessage()")
	return nil
}

func (a1 *asyncProducerImpl) SendMessage(message sarama.ProducerMessage) error {
	return a1.PublishMessage(message)
}

func (a1 *asyncProducerImpl) Stop() {

	logger := getLoggerWithName(_packageNameAsyncProducer + ":Stop()")
	fields := map[string]interface{}{"config": a1.config.string(), "_terminationDelay": _terminationDelay.String()}
	logger.Info().Fields(fields).Msg("Stopping AsyncProducer using `stopChan`")

	if a1.hasStopped.Load() {
		logger.Warn().Fields(fields).Msg("WARNING! AsyncProducer is already closed.")
		return
	}

	a1.hasStopped.Store(true)
	close(a1.stopChan)
	_ = _wr.Close()
	if a1.principal != "" {
		deleteTmpCerts(a1.principal)
	}

	time.Sleep(_terminationDelay)
	logger.Info().Fields(fields).Msg("Finished termination delay")
}

// Start - Ref: https://pkg.go.dev/github.com/Shopify/sarama#AsyncProducer
func (a1 *asyncProducerImpl) Start() {

	logger := getLoggerWithName(_packageNameAsyncProducer + ":Start()")
	// Temporary set log level to INFO
	logger = logger.Level(zerolog.InfoLevel)
	fields := map[string]interface{}{"config": a1.config.string(), "_terminationDelay": _terminationDelay.String()}

	if a1.hasStarted.Load() {
		logger.Info().Fields(fields).Msg("AsyncProducer has already been started!")
		return
	}
	a1.hasStarted.Store(true)
	logger.Info().Fields(fields).Msg("Starting AsyncProducer.")
	// Reset log level
	logger = logger.Level(_logLevel)

	go func() {
		ticker := time.NewTicker(a1.durationToResetCounter)

		defer func() {
			logger.Info().Fields(fields).Msg("Closing AsyncProducer.")
			err := a1.asyncProducer.Close()
			if err != nil {
				logger.Error().Fields(fields).Err(err).Msg("error closing AsyncProducer")
			}
			logger.Info().Fields(fields).Msg("Stopping AsyncProducer ticker.")
			ticker.Stop()
		}()

		keepRunning := true
		for keepRunning {
			select {
			case msg := <-a1.asyncProducer.Successes():
				a1.successCount.Add(1)
				encodedKey := getProducerEncodedKey(msg)
				logger.Info().
					Str("topic", msg.Topic).
					Str("key", encodedKey).
					Msg("SUCCESS producing message")
			case err := <-a1.asyncProducer.Errors():
				a1.funcErrorHandling(err)
				logger.Error().Fields(fields).Err(err.Err).Msg("ERROR producing message")
				a1.errorCount.Add(1)
			case <-ticker.C:
				enqueuedCountStr := _printer.Sprintf(_formatDigit, a1.enqueuedCount.Load())
				successCountStr := _printer.Sprintf(_formatDigit, a1.successCount.Load())
				errorCountStr := _printer.Sprintf(_formatDigit, a1.errorCount.Load())
				logger.Info().
					Str("durationToReset", a1.durationToResetCounter.String()).
					Str("enqueuedCount", enqueuedCountStr).
					Str("successCount", successCountStr).
					Str("errorCount", errorCountStr).Msg("resetting ticker")
				a1.resetCount()
			case <-a1.stopChan:
				keepRunning = false
			}
		}

		logger.Info().Fields(fields).Msg("AsyncProducer shutting down.")
	}()
}

func (a1 *asyncProducerImpl) resetCount() {
	a1.enqueuedCount.Store(0)
	a1.successCount.Store(0)
	a1.errorCount.Store(0)
}

func getProducerEncodedKey(msg *sarama.ProducerMessage) string {
	now := time.Now().UTC()
	nowStr := now.Format(timeisoparser.ISO8601Millis)
	if msg == nil || msg.Key == nil {
		return nowStr
	}
	encodedKey, err := msg.Key.Encode()
	if err != nil {
		return nowStr
	}
	return string(encodedKey)
}
