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
type ConsumerGroupSaslTestSuite struct {
	suite.Suite
	ctxBg             context.Context
	config            ConsumerGroupConfigSasl
	mockConsumerGroup *mockConsumerGroup
	mockProcessor     *mockProcessor
}

func (s *ConsumerGroupSaslTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
	s.mockConsumerGroup = newMockConsumerGroup()
	duration := 50 * time.Millisecond
	s.config = ConsumerGroupConfigSasl{
		Principal:              "username@realm.net",
		Brokers:                []string{"broker1:9094", "broker2:9094"},
		KerbKeytab:             []byte("kerbKeytab"),
		ConsumerGroupId:        "consumerGroupId",
		Topics:                 []string{"topic1", "topic2"},
		DurationToResetCounter: &duration,
	}
	s.mockProcessor = newMockProcessor()
	_saramaNewConsumerGroup = func(addrs []string, groupID string, config *sarama.Config) (sarama.ConsumerGroup, error) {
		return s.mockConsumerGroup, nil
	}
}

func (s *ConsumerGroupSaslTestSuite) TearDownTest() {
	s.resetMonkeyPatching()
	s.mockConsumerGroup.stop()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestConsumerGroupSaslTestSuite(t *testing.T) {
	suite.Run(t, new(ConsumerGroupSaslTestSuite))
}

// #endregion

// ############################################################################
// #region TESTS ARE BELOW
// ############################################################################

func (s *ConsumerGroupSaslTestSuite) Test_GivenProperConsumer_ThenConsumeOk() {
	// -- ARRANGE --
	sutConsumer := NewConsumerWrapperSaslSslAutoStart(s.config, s.mockProcessor)
	defer sutConsumer.Stop()
	// -- ACT --
	s.mockConsumerGroup.errorChan <- errForTests
	time.Sleep(100 * time.Millisecond)
	// -- ASSERT --
	assert.Equal(s.T(), 1, s.mockConsumerGroup.mpfConsume.GetCount())
}

func (s *ConsumerGroupSaslTestSuite) Test_GivenConfigValidationError_ThenPanic() {
	// -- ARRANGE --
	s.config.Brokers = nil
	assert.Panics(s.T(), func() {
		NewConsumerWrapperSaslSslAutoStart(s.config, s.mockProcessor)
	})
	// -- ACT --
	// -- ASSERT --
}

func (s *ConsumerGroupSaslTestSuite) Test_GivenNewConsumerGroupError_ThenPanic() {
	// -- ARRANGE --
	_saramaNewConsumerGroup = sarama.NewConsumerGroup
	assert.Panics(s.T(), func() {
		NewConsumerWrapperSaslSslAutoStart(s.config, s.mockProcessor)
	})
	// -- ACT --
	// -- ASSERT --
}

// ############################################################################
// #region TEST HELPERS
// ############################################################################

func (s *ConsumerGroupSaslTestSuite) resetMonkeyPatching() {
}

// #endregion