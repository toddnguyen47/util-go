package sarama_msk_wrapper

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// /--------------------------------------------------------------------------\
// #region SETUP
// ----------------------------------------------------------------------------

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ConsumerGroupHandlerBatchTestSuite struct {
	suite.Suite
	ctxBg                  context.Context
	mockBatchProcessor     *mockBatchProcessor
	sutImpl                consumerGroupHandlerWithChan
	timeout                time.Duration
	batchSize              int
	mockConsumerGroupClaim *mockConsumerGroupClaimStruct
	mockSess               *mockConsumerGroupSession
}

func (s *ConsumerGroupHandlerBatchTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
	SetLogLevel("INFO")
	s.mockBatchProcessor = newMockBatchProcessor()
	s.timeout = 100 * time.Millisecond
	s.batchSize = 5
	s.mockConsumerGroupClaim = newMockConsumerGroupClaimStruct()
	s.mockSess = newMockConsumerGroupSession()

	// Set up sut (situation under test) now
	s.sutImpl = newConsumerGroupHandlerBatch(s.mockBatchProcessor, uint(s.batchSize), s.timeout)
	err := s.sutImpl.Setup(s.mockSess)
	assert.Nil(s.T(), err)
}

func (s *ConsumerGroupHandlerBatchTestSuite) TearDownTest() {
	s.resetMonkeyPatching()
	err := s.sutImpl.Cleanup(s.mockSess)
	assert.Nil(s.T(), err)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestConsumerGroupHandlerBatchTestSuite(t *testing.T) {
	suite.Run(t, new(ConsumerGroupHandlerBatchTestSuite))
}

// ----------------------------------------------------------------------------
// #endregion SETUP
// \--------------------------------------------------------------------------/

// /--------------------------------------------------------------------------\
// #region TESTS ARE BELOW
// ----------------------------------------------------------------------------

func (s *ConsumerGroupHandlerBatchTestSuite) Test_GivenSuccessfulBatchProcessor_ThenReturnNil() {
	// -- ARRANGE --
	var wg sync.WaitGroup
	numMessages := s.batchSize * 5
	wg.Add(numMessages)
	go s.mockSendingMessagesToGroupClaim(&wg, numMessages, true, 0*time.Millisecond)
	// -- ACT --
	s.sutImpl.MarkNotReady()
	err := s.sutImpl.ConsumeClaim(s.mockSess, s.mockConsumerGroupClaim)
	wg.Wait()
	// -- ASSERT --
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), s.batchSize, s.mockBatchProcessor.mpfBatch.GetCount())
	assert.Equal(s.T(), numMessages, s.mockBatchProcessor.mpfProcess.GetCount())
	assert.Equal(s.T(), numMessages, s.mockSess.mpfMarkMessage.GetCount())
}

func (s *ConsumerGroupHandlerBatchTestSuite) Test_GivenSuccessfulBatchProcessorWithCancellingContext_ThenReturnNil() {
	// -- ARRANGE --
	var wg sync.WaitGroup
	numMessages := s.batchSize * 5
	wg.Add(numMessages)
	go s.mockSendingMessagesToGroupClaim(&wg, numMessages, false, 0*time.Millisecond)
	// -- ACT --
	go func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("cancelling session")
		s.mockSess.cancel()
	}()
	err := s.sutImpl.ConsumeClaim(s.mockSess, s.mockConsumerGroupClaim)
	wg.Wait()
	// -- ASSERT --
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), s.batchSize, s.mockBatchProcessor.mpfBatch.GetCount())
	assert.Equal(s.T(), numMessages, s.mockBatchProcessor.mpfProcess.GetCount())
	assert.Equal(s.T(), numMessages, s.mockSess.mpfMarkMessage.GetCount())
	close(s.mockConsumerGroupClaim.chanConsumerMessage)
}

func (s *ConsumerGroupHandlerBatchTestSuite) Test_GivenTimedOut_ThenProcessCurrentBatch() {
	// -- ARRANGE --
	var wg sync.WaitGroup
	numMessages := s.batchSize * 5
	wg.Add(numMessages)
	go s.mockSendingMessagesToGroupClaim(&wg, numMessages, true, s.timeout>>1)
	// -- ACT --
	err := s.sutImpl.ConsumeClaim(s.mockSess, s.mockConsumerGroupClaim)
	wg.Wait()
	// -- ASSERT --
	assert.Nil(s.T(), err)
	assert.Greater(s.T(), s.mockBatchProcessor.mpfBatch.GetCount(), s.batchSize)
	assert.Equal(s.T(), numMessages, s.mockBatchProcessor.mpfProcess.GetCount())
	assert.Equal(s.T(), numMessages, s.mockSess.mpfMarkMessage.GetCount())
}

func (s *ConsumerGroupHandlerBatchTestSuite) Test_GivenTwoFailedMessages_ThenDoNotMarkThemAsProcessed() {
	// -- ARRANGE --
	s.mockBatchProcessor.mpfProcess.SetCode("FF")
	var wg sync.WaitGroup
	numMessages := s.batchSize * 5
	wg.Add(numMessages)
	go s.mockSendingMessagesToGroupClaim(&wg, numMessages, true, 0*time.Millisecond)
	// -- ACT --
	err := s.sutImpl.ConsumeClaim(s.mockSess, s.mockConsumerGroupClaim)
	wg.Wait()
	// -- ASSERT --
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), s.batchSize, s.mockBatchProcessor.mpfBatch.GetCount())
	assert.Equal(s.T(), numMessages, s.mockBatchProcessor.mpfProcess.GetCount())
	assert.Equal(s.T(), numMessages-2, s.mockSess.mpfMarkMessage.GetCount())
}

// ----------------------------------------------------------------------------
// #endregion TESTS
// \--------------------------------------------------------------------------/

// /--------------------------------------------------------------------------\
// #region TEST HELPERS
// ----------------------------------------------------------------------------

func (s *ConsumerGroupHandlerBatchTestSuite) resetMonkeyPatching() {
}

// mockSendingMessagesToGroupClaim - call with a go function
func (s *ConsumerGroupHandlerBatchTestSuite) mockSendingMessagesToGroupClaim(
	wg *sync.WaitGroup, numMessages int, doCloseChannel bool, timeDelay time.Duration) {
	for i := 0; i < numMessages; i++ {
		s.mockConsumerGroupClaim.chanConsumerMessage <- &sarama.ConsumerMessage{
			Key:   []byte(fmt.Sprintf("key%d", i)),
			Value: []byte(fmt.Sprintf("value%d", i)),
			Topic: "myTopic",
		}
		time.Sleep(timeDelay)
		wg.Done()
	}
	if doCloseChannel {
		close(s.mockConsumerGroupClaim.chanConsumerMessage)
	}
}

// ----------------------------------------------------------------------------
// #region TEST HELPERS
// \--------------------------------------------------------------------------/
