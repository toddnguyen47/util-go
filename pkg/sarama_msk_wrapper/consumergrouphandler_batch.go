package sarama_msk_wrapper

import (
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
	_msgErrorProcessingBatch = "error processing batched consumed messages"

	_strBatchImpl    = "myConsumerGroupHandlerBatchImpl"
	_strBatchSize    = "batchSize"
	_strBatchTimeout = "batchTimeout"
)

// /##########################################################\
// #region myConsumerGroupHandlerBatchImpl
// ############################################################

// myConsumerGroupHandlerBatchImpl - Ref: https://pkg.go.dev/github.com/Shopify/sarama#example-ConsumerGroup
type myConsumerGroupHandlerBatchImpl struct {
	batchProcessor ConsumedBatchOfMessagesProcessor
	batchSize      uint
	batchTimeout   time.Duration
	ticker         *time.Ticker
	mutex          sync.Mutex
	batch          []sarama.ConsumerMessage
}

func newConsumerGroupHandlerBatch( // NOSONAR - need lots of parameters
	batchProcessor ConsumedBatchOfMessagesProcessor,
	batchSize uint,
	batchTimeout time.Duration) sarama.ConsumerGroupHandler {

	ticker := time.NewTicker(batchTimeout)

	m := myConsumerGroupHandlerBatchImpl{
		batchProcessor: batchProcessor,
		batchSize:      batchSize,
		batchTimeout:   batchTimeout,
		ticker:         ticker,
		mutex:          sync.Mutex{},
		batch:          make([]sarama.ConsumerMessage, 0),
	}
	return &m
}

func (i1 *myConsumerGroupHandlerBatchImpl) Setup(sess sarama.ConsumerGroupSession) error {

	logger := getLoggerWithName(_strBatchImpl + ":Setup()")
	logger.Info().Str("memberId", sess.MemberID()).
		Int32("generationId", sess.GenerationID()).
		Uint(_strBatchSize, i1.batchSize).
		Str(_strBatchTimeout, i1.batchTimeout.String()).
		Msg("In SetUp() for the following fields")

	return nil
}

func (i1 *myConsumerGroupHandlerBatchImpl) Cleanup(sess sarama.ConsumerGroupSession) error {

	logger := getLoggerWithName(_strBatchImpl + ":Cleanup()")
	logger.Info().Str("memberId", sess.MemberID()).
		Int32("generationId", sess.GenerationID()).
		Uint(_strBatchSize, i1.batchSize).
		Str(_strBatchTimeout, i1.batchTimeout.String()).
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

	logger := getLoggerWithName(_strBatchImpl + ":ConsumeClaim()")
	fields := map[string]interface{}{
		"memberId":       sess.MemberID(),
		"generationId":   sess.GenerationID(),
		"topic":          claim.Topic(),
		"partition":      claim.Partition(),
		"initialOffset":  claim.InitialOffset(),
		_strBatchSize:    i1.batchSize,
		_strBatchTimeout: i1.batchTimeout.String(),
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
			if len(i1.batch) >= int(i1.batchSize) {
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
		logger := getLoggerWithName(_strBatchImpl + ":processBatch()")
		fields := map[string]interface{}{
			"memberId":         sess.MemberID(),
			"generationId":     sess.GenerationID(),
			"currentBatchSize": lenBatch,
			_strBatchSize:      i1.batchSize,
			_strBatchTimeout:   i1.batchTimeout.String(),
		}
		logger.Info().Fields(fields).Msg("INIT processing batch of consumer messages")
		successfullyConsumedMessages, err := i1.batchProcessor.ProcessConsumedBatchOfMessages(i1.batch)
		fields["lenSuccessfullyConsumedMessages"] = len(successfullyConsumedMessages)
		logger.Info().Fields(fields).Msg("END processing batch of consumer messages")
		if err != nil {
			logger.Error().Err(err).Fields(fields).Msg(_msgErrorProcessingBatch)
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
