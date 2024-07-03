package sarama_msk_wrapper

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// ------------------------------------------------------------
// #region SETUP
// ------------------------------------------------------------

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ConsumerGroupDisabledTestSuite struct {
	suite.Suite
	ctxBg context.Context
}

func (s *ConsumerGroupDisabledTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
}

func (s *ConsumerGroupDisabledTestSuite) TearDownTest() {
	s.resetMonkeyPatching()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestConsumerGroupDisabledTestSuite(t *testing.T) {
	suite.Run(t, new(ConsumerGroupDisabledTestSuite))
}

// #endregion

// ------------------------------------------------------------
// #region TESTS ARE BELOW
// ------------------------------------------------------------

func (s *ConsumerGroupDisabledTestSuite) Test_GivenDisabledConsumerWrapper_ThenDoNotConsume() {
	// -- GIVEN --
	sutDisabled := NewDisabledConsumerWrapper()
	sutDisabled.Start()
	assert.False(s.T(), sutDisabled.HasStopped())
	// -- WHEN --
	sutDisabled.SetErrorHandlingFunction(func(err error) {})
	// -- THEN --
	assert.NotNil(s.T(), sutDisabled.GetConsumerGroup())
	assert.Equal(s.T(), 0, sutDisabled.GetErrorCount())
	sutDisabled.Stop()
	assert.True(s.T(), sutDisabled.HasStopped())
	sutDisabled.Stop()
	assert.True(s.T(), sutDisabled.HasStopped())
}

// ------------------------------------------------------------
// #region TEST HELPERS
// ------------------------------------------------------------

func (s *ConsumerGroupDisabledTestSuite) resetMonkeyPatching() {
}

// #endregion
