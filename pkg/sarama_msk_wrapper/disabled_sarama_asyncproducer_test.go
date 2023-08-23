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
type DisabledSaramaAsyncProducerTestSuite struct {
	suite.Suite
	ctxBg context.Context
}

func (s *DisabledSaramaAsyncProducerTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
}

func (s *DisabledSaramaAsyncProducerTestSuite) TearDownTest() {
	s.resetMonkeyPatching()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestDisabledSaramaAsyncProducerTestSuite(t *testing.T) {
	suite.Run(t, new(DisabledSaramaAsyncProducerTestSuite))
}

// #endregion

// ############################################################################
// #region TESTS ARE BELOW
// ############################################################################

func (s *DisabledSaramaAsyncProducerTestSuite) Test_GivenDisabledAsyncProducer_ThenEveryFuncIsOk() {
	// -- ARRANGE --
	sutDisabled := NewDisabledSaramaAsyncProducer()
	// -- ACT --
	// -- ASSERT --
	assert.NotNil(s.T(), sutDisabled.Input())
	assert.NotNil(s.T(), sutDisabled.Successes())
	assert.NotNil(s.T(), sutDisabled.Errors())
	assert.False(s.T(), sutDisabled.IsTransactional())
	assert.Equal(s.T(), sarama.ProducerTxnStatusFlag(0), sutDisabled.TxnStatus())
	err := sutDisabled.BeginTxn()
	assert.Nil(s.T(), err)
	err = sutDisabled.CommitTxn()
	assert.Nil(s.T(), err)
	err = sutDisabled.AbortTxn()
	assert.Nil(s.T(), err)
	err = sutDisabled.AddOffsetsToTxn(nil, "")
	assert.Nil(s.T(), err)
	err = sutDisabled.AddMessageToTxn(nil, "", nil)
	assert.Nil(s.T(), err)
	sutDisabled.AsyncClose()
	time.Sleep(100 * time.Millisecond)
	err = sutDisabled.Close()
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), sutDisabled.Input())
}

// ############################################################################
// #region TEST HELPERS
// ############################################################################

func (s *DisabledSaramaAsyncProducerTestSuite) resetMonkeyPatching() {
}

// #endregion
