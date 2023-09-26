package sarama_msk_wrapper

import (
	"fmt"
	"sync/atomic"

	"github.com/toddnguyen47/util-go/pkg/sarama_msk_wrapper/saramaconfig"
)

func NewConsumerWrapperSaslSsl(config ConsumerGroupConfigSasl, processor ConsumedMessageProcessor) ConsumerWrapper {

	logger := getLoggerWithName(_packageNameConsumerGroup + ":NewConsumerWrapperSaslSsl()")
	err := config.validate()
	if err != nil {
		wrappedErr := fmt.Errorf("not all required fields are passed into NewConsumerWrapper | err: %w", err)
		logger.Error().Err(wrappedErr).Send()
		panic(wrappedErr)
	}

	saramaConfig := saramaconfig.GetSaramaConfigSasl(config.Principal, config.KerbKeytab,
		config.KerbConf, config.SslCert)
	saramaConfig.Consumer.Return.Errors = true

	// Start a new consumer group
	consumerGroup, err := _saramaNewConsumerGroup(config.Brokers, config.ConsumerGroupId, saramaConfig)
	if err != nil {
		wrappedErr := fmt.Errorf("error creating a new ConsumerGroup | err = %w", err)
		logger.Error().Err(wrappedErr).Send()
		panic(wrappedErr)
	}

	handlerWrapper := newConsumerGroupHandlerWrapper(processor)

	impl := consumerWrapperImpl{
		config:                      &config,
		consumerGroup:               consumerGroup,
		funcMetricErrorConsuming:    func() {},
		funcErrorHandling:           func(err error) {},
		consumerGroupHandlerWrapper: handlerWrapper,
		hasStarted:                  atomic.Bool{},
		hasStopped:                  atomic.Bool{},
		stopChan:                    make(chan struct{}, 1),
		errorCount:                  atomic.Uint32{},
		topics:                      config.Topics,
		principal:                   config.Principal,
		durationToResetCounter:      DefaultTimerResetTime,
	}

	if config.DurationToResetCounter != nil {
		impl.durationToResetCounter = *config.DurationToResetCounter
	}
	impl.hasStopped.Store(false)
	return &impl
}

func NewConsumerWrapperSaslSslAutoStart(config ConsumerGroupConfigSasl,
	processor ConsumedMessageProcessor) ConsumerWrapper {

	consumerWrapper := NewConsumerWrapperSaslSsl(config, processor)
	consumerWrapper.Start()
	return consumerWrapper
}
