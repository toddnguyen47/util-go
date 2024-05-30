package sarama_msk_wrapper

import "github.com/IBM/sarama"

type disabledConsumerWrapper struct {
	consumerGroup sarama.ConsumerGroup
	hasClosed     bool
}

func (dc *disabledConsumerWrapper) Start() {
	logger := getLoggerWithName(_packageNameConsumerGroupDisabled + ":Start()")
	logger.Warn().Msg("WARNING! Starting a disabled ConsumerGroup")
}

func (dc *disabledConsumerWrapper) Stop() {
	if dc.hasClosed {
		return
	}
	dc.hasClosed = true
	logger := getLoggerWithName(_packageNameConsumerGroupDisabled + ":Stop()")
	logger.Info().Msg("Stopping disabled consumer wrapper.")
}

func (dc *disabledConsumerWrapper) HasStopped() bool {
	return dc.hasClosed
}

func (dc *disabledConsumerWrapper) GetConsumerGroup() sarama.ConsumerGroup { return dc.consumerGroup }

func (dc *disabledConsumerWrapper) GetErrorCount() int { return 0 }

func (dc *disabledConsumerWrapper) SetErrorHandlingFunction(_ func(err error)) {}
