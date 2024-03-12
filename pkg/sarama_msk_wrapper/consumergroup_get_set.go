package sarama_msk_wrapper

import "github.com/IBM/sarama"

func (c1 *consumerWrapperImpl) GetConsumerGroup() sarama.ConsumerGroup { return c1.consumerGroup }

func (c1 *consumerWrapperImpl) HasStopped() bool {
	return c1.hasStopped.Load()
}

func (c1 *consumerWrapperImpl) GetErrorCount() int {
	num := c1.errorCount.Load()
	return int(num)
}

func (c1 *consumerWrapperImpl) SetErrorHandlingFunction(myFunc func(err error)) {
	c1.funcErrorHandling = myFunc
}
