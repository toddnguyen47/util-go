package sarama_msk_wrapper

import (
	"context"
	"testing"

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
type AsyncProducerDisabledTestSuite struct {
	suite.Suite
	ctxBg   context.Context
	message sarama.ProducerMessage
}

func (s *AsyncProducerDisabledTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
	s.message = sarama.ProducerMessage{
		Topic: "myTopic",
		Key:   errEncoder(1),
		Value: errEncoder(2),
	}
}

func (s *AsyncProducerDisabledTestSuite) TearDownTest() {
	s.resetMonkeyPatching()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestAsyncProducerDisabledTestSuite(t *testing.T) {
	suite.Run(t, new(AsyncProducerDisabledTestSuite))
}

// #endregion

// ------------------------------------------------------------
// #region TESTS ARE BELOW
// ------------------------------------------------------------

func (s *AsyncProducerDisabledTestSuite) Test_GivenDisabledAsyncProducer_ThenPublishNothing() {
	// -- GIVEN --
	sutDisabledAsyncProducer := NewDisabledAsyncProducer()
	sutDisabledAsyncProducer.Start()
	assert.False(s.T(), sutDisabledAsyncProducer.HasClosed())
	sutDisabledAsyncProducer.SetErrorHandlingFunction(func(err *sarama.ProducerError) {})
	// -- WHEN --
	err := sutDisabledAsyncProducer.PublishMessage(s.message)
	// -- THEN --
	assert.Nil(s.T(), err)
	err = sutDisabledAsyncProducer.SendMessage(s.message)
	assert.Nil(s.T(), err)
	sutDisabledAsyncProducer.Stop()
	assert.True(s.T(), sutDisabledAsyncProducer.HasClosed())
	sutDisabledAsyncProducer.Stop()
	assert.True(s.T(), sutDisabledAsyncProducer.HasClosed())
	assert.NotNil(s.T(), sutDisabledAsyncProducer.GetAsyncProducer())
	assert.Equal(s.T(), 0, sutDisabledAsyncProducer.GetEnqueuedCount())
	assert.Equal(s.T(), 0, sutDisabledAsyncProducer.GetSuccessCount())
	assert.Equal(s.T(), 0, sutDisabledAsyncProducer.GetErrorCount())
}

// ------------------------------------------------------------
// #region TEST HELPERS
// ------------------------------------------------------------

func (s *AsyncProducerDisabledTestSuite) resetMonkeyPatching() {
}

// #endregion
