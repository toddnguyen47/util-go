package sarama_msk_wrapper

import (
	"github.com/IBM/sarama"
)

func (a1 *asyncProducerImpl) HasClosed() bool { return a1.hasStopped.Load() }

func (a1 *asyncProducerImpl) GetAsyncProducer() sarama.AsyncProducer { return a1.asyncProducer }

func (a1 *asyncProducerImpl) GetEnqueuedCount() int {
	num := a1.enqueuedCount.Load()
	return int(num)
}

func (a1 *asyncProducerImpl) GetSuccessCount() int {
	num := a1.successCount.Load()
	return int(num)
}

func (a1 *asyncProducerImpl) GetErrorCount() int {
	num := a1.errorCount.Load()
	return int(num)
}

func (a1 *asyncProducerImpl) SetErrorHandlingFunction(myFunc func(err *sarama.ProducerError)) {
	a1.funcErrorHandling = myFunc
}
