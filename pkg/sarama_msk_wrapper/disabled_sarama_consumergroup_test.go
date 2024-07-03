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
type DisabledConsumerGroupTestSuite struct {
	suite.Suite
	ctxBg context.Context
}

func (s *DisabledConsumerGroupTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
}

func (s *DisabledConsumerGroupTestSuite) TearDownTest() {
	s.resetMonkeyPatching()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestDisabledConsumerGroupTestSuite(t *testing.T) {
	suite.Run(t, new(DisabledConsumerGroupTestSuite))
}

// #endregion

// ------------------------------------------------------------
// #region TESTS ARE BELOW
// ------------------------------------------------------------

func (s *DisabledConsumerGroupTestSuite) Test_GivenDisabledConsumer_ThenAllFunctionsWorkFine() {
	// -- GIVEN --
	sutDisabled := NewDisabledSaramaConsumerGroup()
	// -- WHEN --
	// -- THEN --
	assert.Nil(s.T(), sutDisabled.Consume(s.ctxBg, []string{"topic1", "topic2"}, nil))
	assert.NotNil(s.T(), sutDisabled.Errors())
	assert.Nil(s.T(), sutDisabled.Close())
	assert.Nil(s.T(), sutDisabled.Close())
	sutDisabled.Pause(nil)
	sutDisabled.Resume(nil)
	sutDisabled.PauseAll()
	sutDisabled.ResumeAll()
}

// ------------------------------------------------------------
// #region TEST HELPERS
// ------------------------------------------------------------

func (s *DisabledConsumerGroupTestSuite) resetMonkeyPatching() {
}

// #endregion
