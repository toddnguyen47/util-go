package sarama_msk_wrapper

import (
	"context"
	"testing"
	"time"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// ############################################################################
// #region SETUP
// ############################################################################

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type SampleConsumerGroupTestSuite struct {
	suite.Suite
	ctxBg         context.Context
	mockProcessor *mockProcessor
}

func (s *SampleConsumerGroupTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
	s.mockProcessor = newMockProcessor()
}

func (s *SampleConsumerGroupTestSuite) TearDownTest() {
	s.resetMonkeyPatching()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSampleConsumerGroupTestSuite(t *testing.T) {
	suite.Run(t, new(SampleConsumerGroupTestSuite))
}

// #endregion

// ############################################################################
// #region TESTS ARE BELOW
// ############################################################################

func (s *SampleConsumerGroupTestSuite) Test_GivenSetup_ThenMakeSureReadyChanIsReady() {
	// -- ARRANGE --
	sutConsumerGroupHandlerWrapper := newConsumerGroupHandlerWrapper(s.mockProcessor)
	s.mockProcessor.mpfProcess.SetCode("F")
	mockSess := newMockConsumerGroupSession()
	// -- ACT --
	// -- ASSERT --
	err := sutConsumerGroupHandlerWrapper.Setup(mockSess)
	assert.Nil(s.T(), err)
	time.Sleep(50 * time.Millisecond)
	err = sutConsumerGroupHandlerWrapper.Cleanup(mockSess)
	assert.Nil(s.T(), err)
	mockClaim := newMockConsumerGroupClaimStruct()
	go func() {
		msg := sarama.ConsumerMessage{
			Key:   []byte("key"),
			Value: []byte("value"),
		}
		mockClaim.chanConsumerMessage <- &msg
		mockClaim.chanConsumerMessage <- &msg
		mockClaim.chanConsumerMessage <- &msg
		time.Sleep(250 * time.Millisecond)
		// close messages channel now
		close(mockClaim.chanConsumerMessage)
	}()
	err = sutConsumerGroupHandlerWrapper.ConsumeClaim(mockSess, mockClaim)
	assert.Nil(s.T(), err)
}

// ############################################################################
// #region TEST HELPERS
// ############################################################################

func (s *SampleConsumerGroupTestSuite) resetMonkeyPatching() {
}

type mockConsumerGroupClaimStruct struct {
	sarama.ConsumerGroupClaim

	chanConsumerMessage chan *sarama.ConsumerMessage
}

func newMockConsumerGroupClaimStruct() *mockConsumerGroupClaimStruct {
	m := mockConsumerGroupClaimStruct{
		chanConsumerMessage: make(chan *sarama.ConsumerMessage),
	}
	return &m
}

func (m *mockConsumerGroupClaimStruct) Topic() string {
	return "topic"
}

func (m *mockConsumerGroupClaimStruct) Partition() int32 {
	return 1
}

func (m *mockConsumerGroupClaimStruct) InitialOffset() int64 {
	return 200
}

func (m *mockConsumerGroupClaimStruct) Messages() <-chan *sarama.ConsumerMessage {
	return m.chanConsumerMessage
}

// #endregion
