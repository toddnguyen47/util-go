package sarama_msk_wrapper

import (
	"strconv"
	"sync/atomic"

	"github.com/IBM/sarama"
)

type disabledAsyncProducerWrapper struct {
	hasClosed     atomic.Bool
	asyncProducer sarama.AsyncProducer
}

func (da *disabledAsyncProducerWrapper) Start() {}

func (da *disabledAsyncProducerWrapper) Stop() {

	logger := getLoggerWithName(_packageNameAsyncProducerDisabled + ":Stop()")
	if da.HasClosed() {
		return
	}
	da.hasClosed.Store(true)
	logger.Info().Msg("Stopping disabled producer.")
}

func (da *disabledAsyncProducerWrapper) PublishMessage(message sarama.ProducerMessage) error {

	logger := getLoggerWithName(_packageNameAsyncProducerDisabled + ":PublishMessage()")
	key, err := message.Key.Encode()
	if err != nil {
		key = []byte(strconv.Itoa(message.Key.Length()))
	}
	value, err := message.Value.Encode()
	if err != nil {
		value = []byte(strconv.Itoa(message.Value.Length()))
	}
	logger.Warn().Str("key", string(key)).Str("value", string(value)).
		Msg("WARNING: AsyncProducer is disabled")
	return nil
}

func (da *disabledAsyncProducerWrapper) SendMessage(message sarama.ProducerMessage) error {
	return da.PublishMessage(message)
}

func (da *disabledAsyncProducerWrapper) HasClosed() bool { return da.hasClosed.Load() }

func (da *disabledAsyncProducerWrapper) GetAsyncProducer() sarama.AsyncProducer {
	return da.asyncProducer
}

func (da *disabledAsyncProducerWrapper) GetEnqueuedCount() int { return 0 }

func (da *disabledAsyncProducerWrapper) GetSuccessCount() int { return 0 }

func (da *disabledAsyncProducerWrapper) GetErrorCount() int { return 0 }

func (da *disabledAsyncProducerWrapper) SetErrorHandlingFunction(_ func(err *sarama.ProducerError)) {}
