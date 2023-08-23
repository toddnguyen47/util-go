package sarama_msk_wrapper

import (
	"context"

	"github.com/IBM/sarama"
)

type disabledConsumerGroup struct {
	errorChan chan error
	hasClosed bool
}

func NewDisabledSaramaConsumerGroup() sarama.ConsumerGroup {
	dcg := disabledConsumerGroup{
		errorChan: make(chan error, 1),
		hasClosed: false,
	}
	return &dcg
}

func (dcg *disabledConsumerGroup) Consume(_ context.Context, _ []string, _ sarama.ConsumerGroupHandler) error {
	return nil
}

func (dcg *disabledConsumerGroup) Errors() <-chan error {
	return dcg.errorChan
}

func (dcg *disabledConsumerGroup) Close() error {
	if dcg.hasClosed {
		return nil
	}
	dcg.hasClosed = true
	close(dcg.errorChan)
	return nil
}

func (dcg *disabledConsumerGroup) Pause(_ map[string][]int32) {}

func (dcg *disabledConsumerGroup) Resume(_ map[string][]int32) {}

func (dcg *disabledConsumerGroup) PauseAll() {}

func (dcg *disabledConsumerGroup) ResumeAll() {}
