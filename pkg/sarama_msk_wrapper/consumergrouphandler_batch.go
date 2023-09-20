package sarama_msk_wrapper

import (
	"strconv"
	"sync"
	"time"

	"github.com/IBM/sarama"
)

type ConsumedBatchOfMessagesProcessor interface {
	// ProcessConsumedBatchOfMessages - process batch of messages.
	// Return (successfully consumed messages, error)
	ProcessConsumedBatchOfMessages(consumedMessages []sarama.ConsumerMessage) ([]sarama.ConsumerMessage, error)
}

const (
	msgErrorProcessingBatch = "error processing batched consumed messages"
	strBatchImpl            = "myConsumerGroupHandlerBatchImpl"
)

// /##########################################################\
// #region myConsumerGroupHandlerBatchImpl
// ############################################################

// myConsumerGroupHandlerBatchImpl - Ref: https://pkg.go.dev/github.com/Shopify/sarama#example-ConsumerGroup
type myConsumerGroupHandlerBatchImpl struct {
	batchProcessor ConsumedBatchOfMessagesProcessor
	batchSize      int
	ticker         *time.Ticker
	mutex          sync.Mutex
	batch          []sarama.ConsumerMessage
}

func newConsumerGroupHandlerBatch( // NOSONAR - need lots of parameters
	batchProcessor ConsumedBatchOfMessagesProcessor,
	batchTimeout time.Duration,
	batchSize int) sarama.ConsumerGroupHandler {

	ticker := time.NewTicker(batchTimeout)

	m := myConsumerGroupHandlerBatchImpl{
		batchProcessor: batchProcessor,
		ticker:         ticker,
		batchSize:      batchSize,
		mutex:          sync.Mutex{},
		batch:          make([]sarama.ConsumerMessage, 0),
	}
	return &m
}

func (i1 *myConsumerGroupHandlerBatchImpl) Setup(sess sarama.ConsumerGroupSession) error {

	logger := getLoggerWithName(strBatchImpl + ":Setup()")
	logger.Info().Str("memberId", sess.MemberID()).
		Int32("generationId", sess.GenerationID()).
		Msg("In SetUp() for the following fields")

	return nil
}

func (i1 *myConsumerGroupHandlerBatchImpl) Cleanup(sess sarama.ConsumerGroupSession) error {

	logger := getLoggerWithName(strBatchImpl + ":Cleanup()")
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
func (i1 *myConsumerGroupHandlerBatchImpl) ConsumeClaim(sess sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim) error {

	logger := getLoggerWithName(strBatchImpl + ":ConsumeClaim()")
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
				i1.processBeforeCleanup(sess)
				return nil
			}
			i1.addToBatch(msg)
			if len(i1.batch) >= i1.batchSize {
				// Only process if batch size is reached
				i1.processBatch(sess)
			}
		case <-i1.ticker.C:
			// Timed out! Process batch regardless of batch size
			i1.processBatch(sess)
		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/IBM/sarama/issues/1192
		case <-sess.Context().Done():
			logger.Info().Msg("session context was declared 'Done'")
			i1.processBeforeCleanup(sess)
			return nil
		}
	}
}

func (i1 *myConsumerGroupHandlerBatchImpl) addToBatch(message *sarama.ConsumerMessage) {
	i1.mutex.Lock()
	defer i1.mutex.Unlock()
	if message != nil {
		i1.batch = append(i1.batch, *message)
	}
}

func (i1 *myConsumerGroupHandlerBatchImpl) processBatch(sess sarama.ConsumerGroupSession) {
	lenBatch := len(i1.batch)
	if lenBatch > 0 {
		i1.mutex.Lock()
		defer i1.mutex.Unlock()
		logger := getLoggerWithName(strBatchImpl + ":processBatch()")
		fields := map[string]interface{}{
			"memberId":         sess.MemberID(),
			"generationId":     string(sess.GenerationID()),
			"currentBatchSize": lenBatch,
		}
		logger.Info().Fields(fields).Msg("INIT processing batch of consumer messages")
		successfullyConsumedMessages, err := i1.batchProcessor.ProcessConsumedBatchOfMessages(i1.batch)
		fields["lenSuccessfullyConsumedMessages"] = len(successfullyConsumedMessages)
		logger.Info().Fields(fields).Msg("END processing batch of consumer messages")
		if err != nil {
			logger.Error().Err(err).Fields(fields).Msg(msgErrorProcessingBatch)
		}
		// Only mark messages as completed after batch processes successfully
		for _, msg := range successfullyConsumedMessages {
			sess.MarkMessage(&msg, _markedMetadata)
		}
		// Reset batch regardless of completion
		i1.batch = make([]sarama.ConsumerMessage, 0)
	}
}

func (i1 *myConsumerGroupHandlerBatchImpl) processBeforeCleanup(sess sarama.ConsumerGroupSession) {
	i1.processBatch(sess)
	i1.ticker.Stop()
}
