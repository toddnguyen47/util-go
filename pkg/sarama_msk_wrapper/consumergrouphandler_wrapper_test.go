package sarama_msk_wrapper

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// ------------------------------------------------------------
// #region SETUP
// ------------------------------------------------------------

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type SampleConsumerGroupTestSuite struct {
	suite.Suite
	ctxBg         context.Context
	mockProcessor *mockProcessor
	mockSess      *mockConsumerGroupSession
	mockClaim     *mockConsumerGroupClaimStruct

	// situation under test
	sutConsumerGroupHandlerWrapper consumerGroupHandlerWithChan
}

func (s *SampleConsumerGroupTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
	s.mockProcessor = newMockProcessor()
	s.mockSess = newMockConsumerGroupSession()
	s.mockClaim = newMockConsumerGroupClaimStruct()
	s.sutConsumerGroupHandlerWrapper = newConsumerGroupHandlerWrapper(s.mockProcessor)
	err := s.sutConsumerGroupHandlerWrapper.Setup(s.mockSess)
	assert.Nil(s.T(), err)
}

func (s *SampleConsumerGroupTestSuite) TearDownTest() {
	s.resetMonkeyPatching()
	err := s.sutConsumerGroupHandlerWrapper.Cleanup(s.mockSess)
	assert.Nil(s.T(), err)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSampleConsumerGroupTestSuite(t *testing.T) {
	suite.Run(t, new(SampleConsumerGroupTestSuite))
}

// #endregion

// ------------------------------------------------------------
// #region TESTS ARE BELOW
// ------------------------------------------------------------

func (s *SampleConsumerGroupTestSuite) Test_GivenOk_ThenConsumeClaimProperly() {
	// -- GIVEN --
	s.mockProcessor.mpfProcess.SetCode("F")
	go func() {
		msg := sarama.ConsumerMessage{
			Key:   []byte("key"),
			Value: []byte("value"),
		}
		s.mockClaim.chanConsumerMessage <- &msg
		s.mockClaim.chanConsumerMessage <- &msg
		s.mockClaim.chanConsumerMessage <- &msg
		time.Sleep(250 * time.Millisecond)
		// close messages channel now
		close(s.mockClaim.chanConsumerMessage)
	}()
	// -- WHEN --
	s.sutConsumerGroupHandlerWrapper.MarkNotReady()
	// -- THEN --
	err := s.sutConsumerGroupHandlerWrapper.ConsumeClaim(s.mockSess, s.mockClaim)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 3, s.mockProcessor.mpfProcess.GetCount())
	assert.Equal(s.T(), 2, s.mockSess.mpfMarkMessage.GetCount())
}

func (s *SampleConsumerGroupTestSuite) Test_GivenContextCancelled_ThenConsumeClaimProperly() {
	// -- GIVEN --
	s.mockProcessor.mpfProcess.SetCode("F")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		msg := sarama.ConsumerMessage{
			Key:   []byte("key"),
			Value: []byte("value"),
		}
		s.mockClaim.chanConsumerMessage <- &msg
		s.mockClaim.chanConsumerMessage <- &msg
		s.mockClaim.chanConsumerMessage <- &msg
		time.Sleep(250 * time.Millisecond)
		wg.Done()
	}()
	go func() {
		time.Sleep(250 * time.Millisecond)
		s.mockSess.cancel()
	}()
	// -- WHEN --
	// -- THEN --
	err := s.sutConsumerGroupHandlerWrapper.ConsumeClaim(s.mockSess, s.mockClaim)
	wg.Wait()
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 3, s.mockProcessor.mpfProcess.GetCount())
	assert.Equal(s.T(), 2, s.mockSess.mpfMarkMessage.GetCount())
	close(s.mockClaim.chanConsumerMessage)
}

// ------------------------------------------------------------
// #region TEST HELPERS
// ------------------------------------------------------------

func (s *SampleConsumerGroupTestSuite) resetMonkeyPatching() {
}

// #endregion
