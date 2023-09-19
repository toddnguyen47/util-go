package sarama_msk_wrapper

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/IBM/sarama"
	"github.com/toddnguyen47/util-go/pkg/sarama_msk_wrapper/saramaconfig"
	"github.com/toddnguyen47/util-go/pkg/startstopper"
)

// ConsumerWrapper - Make a consumerWrapper, call Start() when you're ready to consume, then call Stop()
// to close the consumer group.
// Ref: https://pkg.go.dev/github.com/Shopify/sarama#ConsumerGroup
// Ref: https://github.com/Shopify/sarama/blob/main/examples/consumergroup/main.go#L102
//
// If you want a disabled ConsumerWrapper, call NewDisabledConsumerWrapper(). If you want to re-enable it,
// call Stop() on the disabled ConsumerWrapper, then make a new ConsumerWrapper with NewConsumerWrapperAutoStart().
type ConsumerWrapper interface {
	startstopper.StartStopper

	HasStopped() bool
	GetConsumerGroup() sarama.ConsumerGroup
	GetErrorCount() int
	SetMetricFunctionErrorConsuming(myFunc func())

	// SetErrorHandlingFunction - If you want to do more error handling, set your error handling function here
	SetErrorHandlingFunction(myFunc func(err error))
}

type consumerWrapperImpl struct {
	config                      configInterface
	consumerGroup               sarama.ConsumerGroup
	funcMetricErrorConsuming    func()
	funcErrorHandling           func(err error)
	consumerGroupHandlerWrapper sarama.ConsumerGroupHandler
	hasStarted                  atomic.Bool
	hasStopped                  atomic.Bool
	stopChan                    chan struct{}
	errorCount                  atomic.Uint32
	topics                      []string
	durationToResetCounter      time.Duration
}

// NewConsumerWrapper - Create a ConsumerWrapper. Note that you need to call `Start()` for the consumer to start
// consuming, and call `Stop()` in a defer call or a shutdown loop.
func NewConsumerWrapper( // NOSONAR - need lots of parameters
	config ConsumerGroupConfig,
	processor ConsumedMessageProcessor) ConsumerWrapper {

	logger := getLoggerWithName(_packageNameConsumerGroup + ":NewConsumerWrapper()")
	err := config.validate()
	if err != nil {
		wrappedErr := fmt.Errorf("not all required fields are passed into NewConsumerWrapper | err: %w", err)
		logger.Error().Err(wrappedErr).Send()
		panic(wrappedErr)
	}

	consumerGroup := newConsumerGroupWithKeys(config)
	handlerWrapper := newConsumerGroupHandlerWrapper(processor)

	impl := consumerWrapperImpl{
		config:                      &config,
		consumerGroup:               consumerGroup,
		funcMetricErrorConsuming:    noopFunc,
		funcErrorHandling:           noopFuncError,
		consumerGroupHandlerWrapper: handlerWrapper,
		hasStarted:                  atomic.Bool{},
		hasStopped:                  atomic.Bool{},
		stopChan:                    make(chan struct{}, 1),
		errorCount:                  atomic.Uint32{},
		topics:                      config.Topics,
		durationToResetCounter:      DefaultTimerResetTime,
	}

	if config.DurationToResetCounter != nil {
		impl.durationToResetCounter = *config.DurationToResetCounter
	}
	impl.hasStopped.Store(false)
	return &impl
}

func NewConsumerWrapperAutoStart(config ConsumerGroupConfig, processor ConsumedMessageProcessor) ConsumerWrapper {
	impl := NewConsumerWrapper(config, processor)
	impl.Start()
	return impl
}

func NewConsumerWrapperWithConsumerGroupHandler(
	config ConsumerGroupConfig,
	handler sarama.ConsumerGroupHandler) ConsumerWrapper {

	err := config.validate()
	if err != nil {
		err2 := fmt.Errorf("not all required fields are passed into NewConsumerWrapper | err: %w", err)
		panic(err2)
	}

	consumerGroup := newConsumerGroupWithKeys(config)

	impl := consumerWrapperImpl{
		config:                      &config,
		consumerGroup:               consumerGroup,
		funcMetricErrorConsuming:    noopFunc,
		funcErrorHandling:           noopFuncError,
		consumerGroupHandlerWrapper: handler,
		hasStarted:                  atomic.Bool{},
		hasStopped:                  atomic.Bool{},
		stopChan:                    make(chan struct{}, 1),
		errorCount:                  atomic.Uint32{},
		topics:                      config.Topics,
		durationToResetCounter:      DefaultTimerResetTime,
	}

	if config.DurationToResetCounter != nil {
		impl.durationToResetCounter = *config.DurationToResetCounter
	}
	impl.hasStopped.Store(false)
	return &impl
}

func NewDisabledConsumerWrapper() ConsumerWrapper {
	logger := getLoggerWithName(_packageNameConsumerGroupDisabled + ":NewConsumerWrapper()")
	logger.Warn().Msg("WARNING! Creating a new disabled ConsumerWrapper")
	impl := disabledConsumerWrapper{
		consumerGroup: NewDisabledSaramaConsumerGroup(),
		hasClosed:     false,
	}
	return &impl
}

// newConsumerGroupWithKeys - create a new ConsumerGroupId.
//
// Please follow the example at https://pkg.go.dev/github.com/Shopify/sarama#example-ConsumerGroup.
// Particularly, you MUST read the `Errors()` channel, otherwise there will be a deadlock.
func newConsumerGroupWithKeys( // NOSONAR - need lots of parameters
	config ConsumerGroupConfig) sarama.ConsumerGroup {

	logger := getLoggerWithName(_packageNameConsumerGroup + ":newConsumerGroupWithKeys()")
	saramaConfig := saramaconfig.GetSaramaConfigSsl(config.PubKey, config.PrivateKey)
	saramaConfig.Consumer.Return.Errors = true

	// Start a new consumerGroup
	consumerGroup, err := _saramaNewConsumerGroup(config.Brokers, config.ConsumerGroupId, saramaConfig)
	if err != nil {
		wrappedErr := fmt.Errorf("error creating a new ConsumerGroup | err = %w", err)
		logger.Error().Err(wrappedErr).Send()
		panic(wrappedErr)
	}

	return consumerGroup
}

type ConsumedMessageProcessor interface {
	ProcessConsumedMessage(consumedMessage *sarama.ConsumerMessage) error
}
