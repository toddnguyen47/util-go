package sarama_msk_wrapper

import (
	"context"
	"testing"
	"time"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/toddnguyen47/util-go/pkg/pointerutils"
)

// ############################################################################
// #region SETUP
// ############################################################################

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ConsumerGroupTestSuite struct {
	suite.Suite
	ctxBg              context.Context
	mockConsumerGroup  *mockConsumerGroup
	metricCount        int
	errorList          []error
	topics             []string
	mockProcessor      *mockProcessor
	mockBatchProcessor *mockBatchProcessor
	config             ConsumerGroupConfig
}

func (s *ConsumerGroupTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	SetLogLevel("WARN")
	s.ctxBg = context.Background()
	s.mockConsumerGroup = newMockConsumerGroup()
	s.errorList = make([]error, 0)
	s.topics = []string{"topic1", "topic2"}
	s.mockProcessor = newMockProcessor()
	s.mockBatchProcessor = newMockBatchProcessor()
	s.metricCount = 0
	pubKey, privateKey := getCerts(s.T())
	s.config = ConsumerGroupConfig{
		Common: ConsumerGroupConfigCommon{
			Brokers:         []string{"my-kafka-server:9094", "mykafka-server-2:9094"},
			ConsumerGroupId: "ConsumerGroupId",
			Topics:          s.topics,
			BatchSize:       5,
			BatchTimeout:    pointerutils.PtrDuration(100 * time.Millisecond),
		},
		PubKey:     pubKey,
		PrivateKey: privateKey,
	}
}

func (s *ConsumerGroupTestSuite) TearDownTest() {
	s.resetMonkeyPatching()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestConsumerGroupTestSuite(t *testing.T) {
	suite.Run(t, new(ConsumerGroupTestSuite))
}

// #endregion

// ############################################################################
// #region TESTS ARE BELOW
// ############################################################################

func (s *ConsumerGroupTestSuite) Test_GivenConsumerGroupInitOk_ThenReturnProperObject() {
	// -- ARRANGE --
	SetLogLevel("INFO")
	_saramaNewConsumerGroup = func(addrs []string, groupID string, config *sarama.Config) (sarama.ConsumerGroup, error) {
		return s.mockConsumerGroup, nil
	}
	sleepTime := 100 * time.Millisecond
	// -- ACT --
	sutConsumerWrapper := NewConsumerWrapper(s.config, s.mockProcessor)
	sutConsumerWrapper.SetErrorHandlingFunction(func(err error) {
		s.errorList = append(s.errorList, err)
	})
	sutConsumerWrapper.SetMetricFunctionErrorConsuming(func() {
		s.metricCount += 1
	})
	sutConsumerWrapper.Start()
	// Starting twice on purpose for testing
	sutConsumerWrapper.Start()
	s.mockConsumerGroup.errorChan <- errForTests
	time.Sleep(sleepTime * 5)
	assert.False(s.T(), sutConsumerWrapper.HasStopped())
	defer func() {
		sutConsumerWrapper.Stop()
		sutConsumerWrapper.Stop()
		time.Sleep(sleepTime)
		assert.True(s.T(), sutConsumerWrapper.HasStopped())
	}()
	time.Sleep(sleepTime)
	// -- ASSERT --
	assert.NotNil(s.T(), sutConsumerWrapper)
	assert.Equal(s.T(), 1, sutConsumerWrapper.GetErrorCount())
}

func (s *ConsumerGroupTestSuite) Test_GivenConsumerGroupBatchInitOk_ThenReturnNil() {
	// -- ARRANGE --
	_saramaNewConsumerGroup = func(addrs []string, groupID string, config *sarama.Config) (sarama.ConsumerGroup, error) {
		return s.mockConsumerGroup, nil
	}
	sleepTime := 100 * time.Millisecond
	// -- ACT --
	s.config.Common.BatchSize = 5
	s.config.Common.BatchTimeout = pointerutils.PtrDuration(1 * time.Minute)
	s.config.Common.DurationToResetCounter = pointerutils.PtrDuration(30 * time.Minute)
	sutConsumerWrapper := NewConsumerWrapperBatchAutoStart(s.config, s.mockBatchProcessor)
	// Starting twice on purpose for testing
	sutConsumerWrapper.Start()
	s.mockConsumerGroup.errorChan <- errForTests
	time.Sleep(500 * time.Millisecond)
	assert.False(s.T(), sutConsumerWrapper.HasStopped())
	defer func() {
		sutConsumerWrapper.Stop()
		sutConsumerWrapper.Stop()
		time.Sleep(sleepTime)
		assert.True(s.T(), sutConsumerWrapper.HasStopped())
	}()
	time.Sleep(sleepTime)
	// -- ASSERT --
	assert.NotNil(s.T(), sutConsumerWrapper)
	assert.Equal(s.T(), 1, sutConsumerWrapper.GetErrorCount())
}

func (s *ConsumerGroupTestSuite) Test_GivenConsumerGroupInitOkContextCancelled_ThenReturnProperObject() {
	// -- ARRANGE --
	_saramaNewConsumerGroup = func(addrs []string, groupID string, config *sarama.Config) (sarama.ConsumerGroup, error) {
		return s.mockConsumerGroup, nil
	}
	sleepTime := 100 * time.Millisecond
	// -- ACT --
	sutConsumerWrapper := NewConsumerWrapper(s.config, s.mockProcessor)
	sutConsumerWrapper.Start()
	time.Sleep(500 * time.Millisecond)
	assert.False(s.T(), sutConsumerWrapper.HasStopped())
	defer func() {
		s.mockConsumerGroup.stop()
		sutConsumerWrapper.Stop()
		time.Sleep(sleepTime)
		assert.True(s.T(), sutConsumerWrapper.HasStopped())
	}()
	time.Sleep(sleepTime)
	// -- ASSERT --
	assert.NotNil(s.T(), sutConsumerWrapper)
	assert.NotNil(s.T(), sutConsumerWrapper.GetConsumerGroup())
}

func (s *ConsumerGroupTestSuite) Test_GivenConsumerGroupError_ThenPanic() {
	// -- ARRANGE --
	pubKey, privateKey := getCerts(s.T())
	_saramaNewConsumerGroup = func(addrs []string, groupID string, config *sarama.Config) (sarama.ConsumerGroup, error) {
		return nil, errForTests
	}
	sutConsumerWrapperId := "ConsumerGroupId"
	// -- ACT --
	// -- ASSERT --
	assert.Panics(s.T(), func() {
		config := ConsumerGroupConfig{
			Common: ConsumerGroupConfigCommon{
				ConsumerGroupId: sutConsumerWrapperId,
				Topics:          s.topics,
			},
			PubKey:     pubKey,
			PrivateKey: privateKey,
		}
		newConsumerGroupWithKeys(config)
	})
}

func (s *ConsumerGroupTestSuite) Test_GivenConsumerGroupReturnsErr_ThenReturnEarly() {
	// -- ARRANGE --
	_saramaNewConsumerGroup = func(addrs []string, groupID string, config *sarama.Config) (sarama.ConsumerGroup, error) {
		return s.mockConsumerGroup, nil
	}
	s.mockConsumerGroup.mpfConsume.SetCode("FFP")
	sleepTime := 100 * time.Millisecond
	s.config.Common.DurationToResetCounter = &sleepTime
	// -- ACT --
	sutConsumerWrapper := NewConsumerWrapperAutoStart(s.config, s.mockProcessor)
	sutConsumerWrapper.Start()
	assert.False(s.T(), sutConsumerWrapper.HasStopped())
	defer func() {
		s.mockConsumerGroup.stop()
		sutConsumerWrapper.Stop()
		time.Sleep(sleepTime)
		assert.True(s.T(), sutConsumerWrapper.HasStopped())
	}()
	time.Sleep(sleepTime)
	// -- ASSERT --
	assert.NotNil(s.T(), sutConsumerWrapper)
}

func (s *ConsumerGroupTestSuite) Test_GivenNewConsumerWrapperConfigValidationError_ThenPanic() {
	// -- ARRANGE --
	_saramaNewConsumerGroup = func(addrs []string, groupID string, config *sarama.Config) (sarama.ConsumerGroup, error) {
		return s.mockConsumerGroup, nil
	}
	// -- ACT --
	// -- ASSERT --
	assert.Panics(s.T(), func() {
		config := ConsumerGroupConfig{}
		NewConsumerWrapper(config, s.mockProcessor)
	})
}

func (s *ConsumerGroupTestSuite) Test_GivenCountersReset_ThenCounterIs0() {
	// -- ARRANGE --
	_saramaNewConsumerGroup = func(addrs []string, groupID string, config *sarama.Config) (sarama.ConsumerGroup, error) {
		return s.mockConsumerGroup, nil
	}
	sleepTime := 100 * time.Millisecond
	s.config.Common.DurationToResetCounter = &sleepTime
	// -- ACT --
	sutConsumerWrapper := NewConsumerWrapper(s.config, s.mockProcessor)
	sutConsumerWrapper.Start()
	s.mockConsumerGroup.errorChan <- errForTests
	time.Sleep(500 * time.Millisecond)
	assert.False(s.T(), sutConsumerWrapper.HasStopped())
	defer func() {
		sutConsumerWrapper.Stop()
		time.Sleep(sleepTime)
		assert.True(s.T(), sutConsumerWrapper.HasStopped())
	}()
	time.Sleep(sleepTime)
	// -- ASSERT --
	assert.NotNil(s.T(), sutConsumerWrapper)
	assert.Equal(s.T(), 0, sutConsumerWrapper.GetErrorCount())
}

func (s *ConsumerGroupTestSuite) Test_GivenBatchConfigErrorNoTopics_ThenPanic() {
	// -- ARRANGE --
	_saramaNewConsumerGroup = func(addrs []string, groupID string, config *sarama.Config) (sarama.ConsumerGroup, error) {
		return s.mockConsumerGroup, nil
	}
	// -- ACT --
	// -- ASSERT --
	s.config.Common.BatchSize = 5
	s.config.Common.BatchTimeout = pointerutils.PtrDuration(1 * time.Minute)
	s.config.Common.Topics = make([]string, 0)
	assert.Panics(s.T(), func() {
		NewConsumerWrapperBatchAutoStart(s.config, s.mockBatchProcessor)
	})
}

func (s *ConsumerGroupTestSuite) Test_GivenBatchSizeIsZero_ThenPanic() {
	// -- ARRANGE --
	_saramaNewConsumerGroup = func(addrs []string, groupID string, config *sarama.Config) (sarama.ConsumerGroup, error) {
		return s.mockConsumerGroup, nil
	}
	// -- ACT --
	// -- ASSERT --
	s.config.Common.BatchSize = 0
	s.config.Common.BatchTimeout = pointerutils.PtrDuration(1 * time.Minute)
	assert.Panics(s.T(), func() {
		NewConsumerWrapperBatchAutoStart(s.config, s.mockBatchProcessor)
	})
}

func (s *ConsumerGroupTestSuite) Test_GivenConsumeFailsMoreThan5Times_ThenStop() {
	// -- ARRANGE --
	_saramaNewConsumerGroup = func(addrs []string, groupID string, config *sarama.Config) (sarama.ConsumerGroup, error) {
		return s.mockConsumerGroup, nil
	}
	s.mockConsumerGroup.consumeWaitForStop = false
	SetLogLevel("INFO")
	_terminationDelay = 5 * time.Millisecond
	// -- ACT --
	sutConsumerWrapper := NewConsumerWrapperAutoStart(s.config, s.mockProcessor)
	s.mockConsumerGroup.errorChan <- errForTests
	time.Sleep(100 * time.Millisecond)
	sutConsumerWrapper.Stop()
	// -- ASSERT --
	assert.True(s.T(), sutConsumerWrapper.HasStopped())
	assert.NotNil(s.T(), sutConsumerWrapper)
}

// ############################################################################
// #region TEST HELPERS
// ############################################################################

func (s *ConsumerGroupTestSuite) resetMonkeyPatching() {
	_saramaNewConsumerGroup = sarama.NewConsumerGroup
	_terminationDelay = 50 * time.Millisecond
}
