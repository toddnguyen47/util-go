package sarama_msk_wrapper

import (
	"strconv"

	"github.com/IBM/sarama"
)

const _markedMetadata = "READ_MjAyMy0wNC0yOFQwMDowMDowMC4wMDBa"

type ConsumedMessageProcessor interface {
	ProcessConsumedMessage(consumedMessage *sarama.ConsumerMessage) error
}

// /##########################################################\
// #region myConsumerGroupHandlerImpl
// ############################################################

// myConsumerGroupHandlerImpl - Ref: https://pkg.go.dev/github.com/Shopify/sarama#example-ConsumerGroup
type myConsumerGroupHandlerImpl struct {
	processor ConsumedMessageProcessor
}

func newConsumerGroupHandlerWrapper(processor ConsumedMessageProcessor) sarama.ConsumerGroupHandler {
	m := myConsumerGroupHandlerImpl{
		processor: processor,
	}
	return &m
}

func (i1 *myConsumerGroupHandlerImpl) Setup(sess sarama.ConsumerGroupSession) error {

	logger := getLoggerWithName(_packageNameConsumerGroup + ":Setup()")
	logger.Info().Str("memberId", sess.MemberID()).
		Int32("generationId", sess.GenerationID()).
		Msg("In SetUp() for the following fields")

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
		"generationId":  string(sess.GenerationID()),
		"topic":         claim.Topic(),
		"partition":     string(claim.Partition()),
		"initialOffset": strconv.FormatInt(claim.InitialOffset(), 10),
	}

	logger.Info().Fields(fields).Msg("Started ConsumeClaim")

	for {
		select {
		case msg, ok := <-claim.Messages():
			if !ok {
				logger.Info().Msg("message channel was closed")
				return nil
			}
			sess.MarkMessage(msg, _markedMetadata)
			err := i1.processor.ProcessConsumedMessage(msg)
			if err != nil {
				logger.Error().Err(err).Fields(fields).Msg("error processing consumed message")
				continue
			}
		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/IBM/sarama/issues/1192
		case <-sess.Context().Done():
			logger.Info().Msg("session context was declared 'Done'")
			return nil
		}
	}
}
