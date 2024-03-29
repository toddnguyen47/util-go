package sarama_msk_wrapper

import (
	"github.com/IBM/sarama"
)

const _markedMetadata = "READ_MjAyMy0wNC0yOFQwMDowMDowMC4wMDBa"

type ConsumedMessageProcessor interface {
	ProcessConsumedMessage(consumedMessage *sarama.ConsumerMessage) error
}

type consumerGroupHandlerWithChan interface {
	sarama.ConsumerGroupHandler
	ReadyChan() <-chan struct{}
	MarkNotReady()
}

// /##########################################################\
// #region myConsumerGroupHandlerImpl
// ############################################################

// myConsumerGroupHandlerImpl - Ref: https://pkg.go.dev/github.com/Shopify/sarama#example-ConsumerGroup
type myConsumerGroupHandlerImpl struct {
	processor ConsumedMessageProcessor
	readyChan chan struct{}
}

func newConsumerGroupHandlerWrapper(processor ConsumedMessageProcessor) consumerGroupHandlerWithChan {
	m := myConsumerGroupHandlerImpl{
		processor: processor,
		readyChan: make(chan struct{}),
	}
	return &m
}

func (i1 *myConsumerGroupHandlerImpl) ReadyChan() <-chan struct{} {
	return i1.readyChan
}

func (i1 *myConsumerGroupHandlerImpl) MarkNotReady() {
	i1.readyChan = make(chan struct{})
}

func (i1 *myConsumerGroupHandlerImpl) Setup(sess sarama.ConsumerGroupSession) error {

	logger := getLoggerWithName(_packageNameConsumerGroup + ":Setup()")
	logger.Info().Str("memberId", sess.MemberID()).
		Int32("generationId", sess.GenerationID()).
		Msg("In SetUp() for the following fields")
	// Mark as ready
	close(i1.readyChan)

	return nil
}

func (i1 *myConsumerGroupHandlerImpl) Cleanup(sess sarama.ConsumerGroupSession) error {

	logger := getLoggerWithName(_packageNameConsumerGroup + ":Cleanup()")
	logger.Info().Str("memberId", sess.MemberID()).
		Int32("generationId", sess.GenerationID()).
		Msg("In Cleanup() for the following fields")
	return nil
}

// ConsumeClaim - Ref: https://github.com/IBM/sarama/blob/main/examples/consumergroup/main.go#L179
//
// NOTE:
// Do not move the code below to a goroutine.
// The `ConsumeClaim` itself is called within a goroutine, see:
// https://github.com/IBM/sarama/blob/main/consumer_group.go#L27-L29
func (i1 *myConsumerGroupHandlerImpl) ConsumeClaim(sess sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim) error {

	logger := getLoggerWithName(_packageNameConsumerGroup + ":ConsumeClaim()")
	fields := map[string]interface{}{
		"memberId":      sess.MemberID(),
		"generationId":  sess.GenerationID(),
		"topic":         claim.Topic(),
		"partition":     claim.Partition(),
		"initialOffset": _printer.Sprintf(_formatDigit, claim.InitialOffset()),
	}

	logger.Info().Fields(fields).Msg("Started ConsumeClaim")

	for {
		select {
		case msg, ok := <-claim.Messages():
			if !ok {
				logger.Info().Msg("message channel was closed")
				return nil
			}
			err := i1.processor.ProcessConsumedMessage(msg)
			if err != nil {
				logger.Error().Err(err).Fields(fields).Msg("error processing consumed message")
				continue
			}
			// Only mark consumed message if it processes successfully
			sess.MarkMessage(msg, _markedMetadata)
		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/IBM/sarama/issues/1192
		case <-sess.Context().Done():
			logger.Info().Msg("session context was declared 'Done'")
			return nil
		}
	}
}
