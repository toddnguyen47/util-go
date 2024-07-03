package sarama_msk_wrapper

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/IBM/sarama"
	"github.com/toddnguyen47/util-go/pkg/sarama_msk_wrapper/saramaconfig"
	"github.com/toddnguyen47/util-go/pkg/startstopper"
)

// AsyncProducerWrapper - wrapper around sarama's `AsyncProducer` to make sending / publishing messages
// easier. Config values are taken from goatt's saramakit, which can be found here:
// https://egbitbucket.dtvops.net/projects/GOATT/repos/saramakit/browse.
//
// You can make a disabled AsyncProducer by calling NewDisabledAsyncProducer(). This disabled AsyncProducer
// will not instantiate nor publish any messages. To get an enabled AsyncProducer, call the Stop() function, then
// make a new AsyncProducer by calling NewAsyncProducerWrapperAutoStart()
type AsyncProducerWrapper interface {

	// StartStopper interface - Stop closes the AsyncProducer and stop it from producing. This will close down the AsyncProducer.
	// This function MUST be called, either with a defer() call or during a shutdown loop.
	startstopper.StartStopper

	// PublishMessage - Publish / Send a message. This function returns an error if the producer is closed.
	// Sample message:
	//
	//	func setUpProducerMessage() sarama.ProducerMessage {
	//		fields := map[string]string{
	//			"map1Key1": "map1Val1",
	//			"map1Key2": "map1Val2",
	//		}
	//		b1, _ := json.Marshal(fields)
	//		msg := sarama.ProducerMessage{
	//			Topic: "myTopic",
	//			Key:   sarama.StringEncoder("myKey"),
	//			Value: sarama.ByteEncoder(b1),
	//		}
	//		return msg
	//	}
	PublishMessage(message sarama.ProducerMessage) error

	// SendMessage - alias for PublishMessage
	SendMessage(message sarama.ProducerMessage) error

	HasClosed() bool
	GetAsyncProducer() sarama.AsyncProducer
	GetEnqueuedCount() int
	GetSuccessCount() int
	GetErrorCount() int

	// SetErrorHandlingFunction - If you want to do more error handling, set your error handling function here
	SetErrorHandlingFunction(myFunc func(err *sarama.ProducerError))
}

type asyncProducerImpl struct {
	config            configInterface
	principal         string
	asyncProducer     sarama.AsyncProducer
	funcErrorHandling func(err *sarama.ProducerError)

	stopChan               chan struct{}
	hasStopped             atomic.Bool
	durationToResetCounter time.Duration
	hasStarted             atomic.Bool

	enqueuedCount atomic.Uint32
	successCount  atomic.Uint32
	errorCount    atomic.Uint32
}

// NewAsyncProducerWrapper - Create a new AsyncProducerWrapper.
// Will default the reset counter timer to DefaultTimerResetTime.
// Note that you will need to call `Start()` for the producers to start producing.
func NewAsyncProducerWrapper( // NOSONAR - need lots of parameters
	config AsyncProducerConfig) AsyncProducerWrapper {

	logger := getLoggerWithName(_packageNameAsyncProducer + ":NewAsyncProducerWrapper()")
	err := config.validate()
	if err != nil {
		wrappedErr := fmt.Errorf("async producer config error | err: %w", err)
		logger.Error().Err(wrappedErr).Send()
		panic(wrappedErr)
	}

	asyncProducer, err := newAsyncProducer(config)
	if err != nil {
		wrappedErr := fmt.Errorf("error creating a new AsyncProducer | err: %w", err)
		logger.Error().Err(wrappedErr).Send()
		panic(wrappedErr)
	}

	impl := asyncProducerImpl{
		config:                 &config,
		asyncProducer:          asyncProducer,
		stopChan:               make(chan struct{}, 1),
		funcErrorHandling:      noopProducerError,
		hasStopped:             atomic.Bool{},
		durationToResetCounter: DefaultTimerResetTime,
	}
	impl.hasStopped.Store(false)
	if config.DurationToResetCounter != nil {
		impl.durationToResetCounter = *config.DurationToResetCounter
	}
	return &impl
}

func NewAsyncProducerWrapperAutoStart(config AsyncProducerConfig) AsyncProducerWrapper {
	impl := NewAsyncProducerWrapper(config)
	impl.Start()
	return impl
}

func NewDisabledAsyncProducer() AsyncProducerWrapper {
	logger := getLoggerWithName(_packageNameAsyncProducerDisabled + ":PublishMessage()")
	logger.Warn().Msg("Creating a new disabled AsyncProducer")
	impl := disabledAsyncProducerWrapper{
		asyncProducer: NewDisabledSaramaAsyncProducer(),
	}
	return &impl
}

// newAsyncProducer - returns (asyncProducer, error)
func newAsyncProducer(config AsyncProducerConfig) (sarama.AsyncProducer, error) {

	saramaConfig := saramaconfig.GetSaramaConfigSsl(config.PubKey, config.PrivateKey)
	saramaConfig.Producer.Return.Successes = true
	saramaConfig.Producer.Return.Errors = true

	asyncProducer, err := _saramaNewAsyncProducer(config.Common.Brokers, saramaConfig)
	return asyncProducer, err
}
