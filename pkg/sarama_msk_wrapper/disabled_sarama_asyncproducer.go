package sarama_msk_wrapper

import "github.com/IBM/sarama"

type disabledAsyncProducer struct {
	chanInput     chan *sarama.ProducerMessage
	chanSuccesses chan *sarama.ProducerMessage
	chanError     chan *sarama.ProducerError
	hasClosed     bool
}

func NewDisabledSaramaAsyncProducer() sarama.AsyncProducer {
	dap := disabledAsyncProducer{
		chanInput:     make(chan *sarama.ProducerMessage, 1),
		chanSuccesses: make(chan *sarama.ProducerMessage, 1),
		chanError:     make(chan *sarama.ProducerError, 1),
		hasClosed:     false,
	}
	return &dap
}

func (dap *disabledAsyncProducer) AsyncClose() {
	go func() {
		_ = dap.Close()
	}()
}

func (dap *disabledAsyncProducer) Close() error {
	if dap.hasClosed {
		return nil
	}
	dap.hasClosed = true
	close(dap.chanInput)
	close(dap.chanSuccesses)
	close(dap.chanError)
	return nil
}

func (dap *disabledAsyncProducer) Input() chan<- *sarama.ProducerMessage {
	return dap.chanInput
}

func (dap *disabledAsyncProducer) Successes() <-chan *sarama.ProducerMessage {
	return dap.chanSuccesses
}

func (dap *disabledAsyncProducer) Errors() <-chan *sarama.ProducerError {
	return dap.chanError
}

func (dap *disabledAsyncProducer) IsTransactional() bool { return false }

func (dap *disabledAsyncProducer) TxnStatus() sarama.ProducerTxnStatusFlag { return 0 }

func (dap *disabledAsyncProducer) BeginTxn() error { return nil }

func (dap *disabledAsyncProducer) CommitTxn() error { return nil }

func (dap *disabledAsyncProducer) AbortTxn() error { return nil }

func (dap *disabledAsyncProducer) AddOffsetsToTxn(
	_ map[string][]*sarama.PartitionOffsetMetadata, _ string) error {
	return nil
}

func (dap *disabledAsyncProducer) AddMessageToTxn(_ *sarama.ConsumerMessage, _ string, _ *string) error {
	return nil
}
