package sarama_msk_wrapper

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/toddnguyen47/util-go/pkg/osutils"
	"github.com/toddnguyen47/util-go/pkg/sarama_msk_wrapper/saramainject"
)

// ------------------------------------------------------------
// #region SETUP
// ------------------------------------------------------------

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type AsyncProducerSaslTestSuite struct {
	suite.Suite
	ctxBg              context.Context
	config             AsyncProducerConfigSasl
	producerMessage    sarama.ProducerMessage
	mockAsyncProducer1 *mockAsyncProducer
}

func (s *AsyncProducerSaslTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
	duration := 50 * time.Millisecond
	s.config = AsyncProducerConfigSasl{
		Common: AsyncProducerConfigCommon{
			Brokers: []string{"broker1:9094", "broker2:9094"},
		},
		Principal:              "username@realm",
		KerbKeytab:             []byte("kerbKeytab"),
		DurationToResetCounter: &duration,
	}
	s.producerMessage = sarama.ProducerMessage{
		Topic: "topic",
		Key:   sarama.StringEncoder("key"),
		Value: sarama.StringEncoder("value"),
	}
	s.mockAsyncProducer1 = newMockAsyncProducer()
	_saramaNewAsyncProducer = func(addrs []string, conf *sarama.Config) (sarama.AsyncProducer, error) {
		return s.mockAsyncProducer1, nil
	}
}

func (s *AsyncProducerSaslTestSuite) TearDownTest() {
	s.resetMonkeyPatching()
	s.mockAsyncProducer1.stop()

	// Remove tmpCertFolder
	tmpCertFolder := saramainject.TmpCertFolder(s.config.Principal)
	files, err := filepath.Glob(tmpCertFolder + "/*")
	assert.Nil(s.T(), err)
	for _, file := range files {
		_ = osutils.RemoveIfExists(file)
	}
	_ = osutils.RemoveIfExists(tmpCertFolder)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestAsyncProducerSaslTestSuite(t *testing.T) {
	suite.Run(t, new(AsyncProducerSaslTestSuite))
}

// #endregion

// ------------------------------------------------------------
// #region TESTS ARE BELOW
// ------------------------------------------------------------

func (s *AsyncProducerSaslTestSuite) Test_GivenEverythingOk_ThenReturnAsyncProducer() {
	// -- GIVEN --
	sutAsyncProducer := NewAsyncProducerSasSslAutoStart(s.config)
	defer sutAsyncProducer.Stop()
	// -- WHEN --
	err := sutAsyncProducer.PublishMessage(s.producerMessage)
	time.Sleep(50 * time.Millisecond)
	// -- THEN --
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 1, getIntFromAtomic(&s.mockAsyncProducer1.inputCount))
}

func (s *AsyncProducerSaslTestSuite) Test_GivenConfigValidationError_ThenPanic() {
	// -- GIVEN --
	s.config.KerbKeytab = nil
	// -- WHEN --
	// -- THEN --
	assert.Panics(s.T(), func() {
		sutAsyncProducer := NewAsyncProducerSasSslAutoStart(s.config)
		defer sutAsyncProducer.Stop()
	})
}

func (s *AsyncProducerSaslTestSuite) Test_GivenConfigCommonValidationError_ThenPanic() {
	// -- GIVEN --
	s.config.Common.Brokers = []string{}
	// -- WHEN --
	// -- THEN --
	assert.Panics(s.T(), func() {
		sutAsyncProducer := NewAsyncProducerSasSslAutoStart(s.config)
		defer sutAsyncProducer.Stop()
	})
}

func (s *AsyncProducerSaslTestSuite) Test_GivenCreatingNewAsyncProducerError_ThenPanic() {
	// -- GIVEN --
	_saramaNewAsyncProducer = sarama.NewAsyncProducer
	// -- WHEN --
	// -- THEN --
	assert.Panics(s.T(), func() {
		sutAsyncProducer := NewAsyncProducerSasSslAutoStart(s.config)
		defer sutAsyncProducer.Stop()
	})
}

func (s *AsyncProducerSaslTestSuite) Test_GivenGlobAndRemoveError_ThenContinue() {
	// -- GIVEN --
	sutAsyncProducer := NewAsyncProducerSasSslAutoStart(s.config)
	// -- WHEN --
	err := sutAsyncProducer.PublishMessage(s.producerMessage)
	time.Sleep(50 * time.Millisecond)
	// -- THEN --
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 1, getIntFromAtomic(&s.mockAsyncProducer1.inputCount))
	filePathFail := true
	_filepathGlob = func(pattern string) (matches []string, err error) {
		if filePathFail {
			filePathFail = false
			return []string{}, errForTests
		}
		return filepath.Glob(pattern)
	}
	removeFail := true
	_osRemove = func(name string) error {
		if removeFail {
			removeFail = false
			return errForTests
		}
		return os.Remove(name)
	}
	sutAsyncProducer.Stop()
}

// ------------------------------------------------------------
// #region TEST HELPERS
// ------------------------------------------------------------

func (s *AsyncProducerSaslTestSuite) resetMonkeyPatching() {
	_saramaNewAsyncProducer = sarama.NewAsyncProducer
	_saramaNewConsumerGroup = sarama.NewConsumerGroup
	_filepathGlob = filepath.Glob
	_osRemove = os.Remove
}

// #endregion
