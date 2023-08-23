package sarama_msk_wrapper

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/toddnguyen47/util-go/pkg/jsonutils"
)

var errForTests = errors.New("errForTests")

// ############################################################################
// #region SETUP
// ############################################################################

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type AsyncProducerTestSuite struct {
	suite.Suite
	ctxBg              context.Context
	mockAsyncProducer1 *mockAsyncProducer
	metricCount        int
	errorList          []error
	config             AsyncProducerConfig
}

func (s *AsyncProducerTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
	s.mockAsyncProducer1 = newMockAsyncProducer()
	s.metricCount = 0
	s.errorList = make([]error, 0)
	pub, pri := getCerts(s.T())
	duration := 5 * time.Second
	s.config = AsyncProducerConfig{
		Brokers:                []string{"my-kafka-server:9094", "mykafka-server-2:9094"},
		PubKey:                 pub,
		PrivateKey:             pri,
		DurationToResetCounter: &duration,
	}
}

func (s *AsyncProducerTestSuite) TearDownTest() {
	s.resetMonkeyPatching()
	s.mockAsyncProducer1.stop()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestAsyncProducerTestSuite(t *testing.T) {
	suite.Run(t, new(AsyncProducerTestSuite))
}

// #endregion

// ############################################################################
// #region TESTS ARE BELOW
// ############################################################################

func (s *AsyncProducerTestSuite) Test_GivenProducerSendsOk_ThenSentOk() {
	// -- ARRANGE --
	s.resetMonkeyPatching()
	_saramaNewAsyncProducer = func(addrs []string, conf *sarama.Config) (sarama.AsyncProducer, error) {
		return s.mockAsyncProducer1, nil
	}
	sleepTimeMillis := 100
	sleepTime := time.Duration(sleepTimeMillis) * time.Millisecond
	sleepTimeHalved := time.Duration(sleepTimeMillis>>1) * time.Millisecond
	s.config.DurationToResetCounter = &sleepTimeHalved
	sutAsyncProducer := NewAsyncProducerWrapper(s.config)
	assert.NotNil(s.T(), sutAsyncProducer.GetAsyncProducer())
	sutAsyncProducer.Start()
	// Start twice on purpose to check
	sutAsyncProducer.Start()
	sutAsyncProducer.SetMetricFunctionErrorProducing(func() {
		s.metricCount += 1
	})
	sutAsyncProducer.SetErrorHandlingFunction(func(err error) {
		s.errorList = append(s.errorList, err)
	})
	msg := s.setUpProducerMessage()
	// -- ACT --
	err := sutAsyncProducer.SendMessage(msg)
	assert.False(s.T(), sutAsyncProducer.HasClosed())
	time.Sleep(sleepTime)
	sutAsyncProducer.Stop()
	s.mockAsyncProducer1.stop()
	// -- ASSERT --
	assert.Nil(s.T(), err)
	assert.True(s.T(), sutAsyncProducer.HasClosed())
	assert.Equalf(s.T(), 1, getIntFromAtomic(&s.mockAsyncProducer1.inputCount), "should be same input count")
	assert.Equal(s.T(), 0, getIntFromAtomic(&s.mockAsyncProducer1.errorCount))
	// These counters have already reset
	assert.Equal(s.T(), 0, sutAsyncProducer.GetEnqueuedCount())
	assert.Equal(s.T(), 0, sutAsyncProducer.GetSuccessCount())
	assert.Equal(s.T(), 0, sutAsyncProducer.GetErrorCount())
	assert.Equal(s.T(), 0, s.metricCount)
}

func (s *AsyncProducerTestSuite) Test_GivenImproperConfig_ThenPanic() {
	// -- ARRANGE --
	s.resetMonkeyPatching()
	assert.Panics(s.T(), func() {
		config := AsyncProducerConfig{}
		_ = NewAsyncProducerWrapper(config)
	})
}

func (s *AsyncProducerTestSuite) Test_GivenStoppedProducer_ThenSendMessageReturnsError() {
	// -- ARRANGE --
	s.resetMonkeyPatching()
	_saramaNewAsyncProducer = func(addrs []string, conf *sarama.Config) (sarama.AsyncProducer, error) {
		return s.mockAsyncProducer1, nil
	}
	sutAsyncProducer := NewAsyncProducerWrapperAutoStart(s.config)
	s.mockAsyncProducer1.closeCode = "FP"
	msg := s.setUpProducerMessage()
	// -- ACT --
	sutAsyncProducer.Stop()
	sutAsyncProducer.Stop()
	err := sutAsyncProducer.PublishMessage(msg)
	// -- ASSERT --
	assert.NotNil(s.T(), err)
}

func (s *AsyncProducerTestSuite) Test_GivenProducerSendsOneErrorOneOk_ThenSentOk() {
	// -- ARRANGE --
	s.resetMonkeyPatching()
	_saramaNewAsyncProducer = func(addrs []string, conf *sarama.Config) (sarama.AsyncProducer, error) {
		return s.mockAsyncProducer1, nil
	}
	sutAsyncProducer := NewAsyncProducerWrapper(s.config)
	sutAsyncProducer.Start()
	s.mockAsyncProducer1.inputErrorCode = "FP"
	msg := s.setUpProducerMessage()
	// -- ACT --
	err := sutAsyncProducer.SendMessage(msg)
	assert.Nil(s.T(), err)
	msg2 := s.setUpProducerMessage()
	msg2.Key = sarama.StringEncoder("key2")
	err = sutAsyncProducer.SendMessage(msg2)
	assert.Nil(s.T(), err)
	time.Sleep(100 * time.Millisecond)
	sutAsyncProducer.Stop()
	s.mockAsyncProducer1.stop()
	// -- ASSERT --=
	assert.True(s.T(), sutAsyncProducer.HasClosed())
	assert.Equalf(s.T(), 2, getIntFromAtomic(&s.mockAsyncProducer1.inputCount), "should be same input count")
	assert.Equal(s.T(), 1, getIntFromAtomic(&s.mockAsyncProducer1.errorCount))
	assert.Equal(s.T(), 2, sutAsyncProducer.GetEnqueuedCount())
	assert.Equal(s.T(), 1, sutAsyncProducer.GetSuccessCount())
	assert.Equal(s.T(), 1, sutAsyncProducer.GetErrorCount())
}

func (s *AsyncProducerTestSuite) Test_GivenGettingCertsError_ThenPanic() {
	// -- ARRANGE --
	s.resetMonkeyPatching()
	_saramaNewAsyncProducer = func(addrs []string, conf *sarama.Config) (sarama.AsyncProducer, error) {
		return s.mockAsyncProducer1, nil
	}
	assert.Panics(s.T(), func() {
		s.config.PubKey = []byte("asdf")
		NewAsyncProducerWrapper(s.config)
	})
}

func (s *AsyncProducerTestSuite) Test_GivenNewAsyncProducerError_ThenPanic() {
	// -- ARRANGE --
	s.resetMonkeyPatching()
	_saramaNewAsyncProducer = func(addrs []string, conf *sarama.Config) (sarama.AsyncProducer, error) {
		return s.mockAsyncProducer1, errForTests
	}
	assert.Panics(s.T(), func() {
		NewAsyncProducerWrapper(s.config)
	})
}

func (s *AsyncProducerTestSuite) Test_GivenMessageKeyParsedError_ThenDoNotPanic() {
	// -- ARRANGE --
	s.resetMonkeyPatching()
	_saramaNewAsyncProducer = func(addrs []string, conf *sarama.Config) (sarama.AsyncProducer, error) {
		return s.mockAsyncProducer1, nil
	}
	sleepTimeMillis := 100
	sleepTime := time.Duration(sleepTimeMillis) * time.Millisecond
	sleepTimeHalved := time.Duration(sleepTimeMillis>>1) * time.Millisecond
	s.config.DurationToResetCounter = &sleepTimeHalved
	sutAsyncProducer := NewAsyncProducerWrapper(s.config)
	assert.NotNil(s.T(), sutAsyncProducer.GetAsyncProducer())
	sutAsyncProducer.Start()
	msg := s.setUpProducerMessage()
	msg.Key = errorEncoder(5)
	// -- ACT --
	err := sutAsyncProducer.PublishMessage(msg)
	assert.False(s.T(), sutAsyncProducer.HasClosed())
	time.Sleep(sleepTime)
	sutAsyncProducer.Stop()
	s.mockAsyncProducer1.stop()
	// -- ASSERT --
	assert.Nil(s.T(), err)
	assert.True(s.T(), sutAsyncProducer.HasClosed())
	assert.Equalf(s.T(), 1, getIntFromAtomic(&s.mockAsyncProducer1.inputCount), "should be same input count")
	assert.Equal(s.T(), 0, getIntFromAtomic(&s.mockAsyncProducer1.errorCount))
	// These counters have already reset
	assert.Equal(s.T(), 0, sutAsyncProducer.GetSuccessCount())
	assert.Equal(s.T(), 0, sutAsyncProducer.GetErrorCount())
}

func (s *AsyncProducerTestSuite) Test_GivenProducerHasNotBeenStartedWhenPublishing_ThenReturnErr() {
	// -- ARRANGE --
	s.resetMonkeyPatching()
	_saramaNewAsyncProducer = func(addrs []string, conf *sarama.Config) (sarama.AsyncProducer, error) {
		return s.mockAsyncProducer1, nil
	}
	var sleepTime = 100 * time.Millisecond
	s.config.DurationToResetCounter = &sleepTime
	sutAsyncProducer := NewAsyncProducerWrapper(s.config)
	assert.NotNil(s.T(), sutAsyncProducer.GetAsyncProducer())
	msg := s.setUpProducerMessage()
	// -- ACT --
	err := sutAsyncProducer.SendMessage(msg)
	assert.False(s.T(), sutAsyncProducer.HasClosed())
	time.Sleep(sleepTime)
	sutAsyncProducer.Stop()
	s.mockAsyncProducer1.stop()
	// -- ASSERT --
	assert.NotNil(s.T(), err)
}

// ############################################################################
// #region TEST HELPERS
// ############################################################################

func (s *AsyncProducerTestSuite) resetMonkeyPatching() {
	_saramaNewAsyncProducer = sarama.NewAsyncProducer
	_terminationDelay = 50 * time.Millisecond
}

func (s *AsyncProducerTestSuite) setUpProducerMessage() sarama.ProducerMessage {
	fields := map[string]string{
		"map1Key1": "map1Val1",
		"map1Key2": "map1Val2",
	}
	b1, _ := jsonutils.MarshalNoEscapeHtml(fields)
	msg := sarama.ProducerMessage{
		Topic: "myTopic",
		Key:   sarama.StringEncoder("myKey"),
		Value: sarama.ByteEncoder(b1),
	}
	return msg
}

type errorEncoder int

func (e errorEncoder) Encode() ([]byte, error) {
	return nil, errForTests
}

func (e errorEncoder) Length() int {
	return 0
}

// #endregion
