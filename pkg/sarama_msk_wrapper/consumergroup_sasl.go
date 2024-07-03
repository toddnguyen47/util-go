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
	consumerGroup, err := _saramaNewConsumerGroup(config.Common.Brokers, config.Common.ConsumerGroupId, saramaConfig)
	if err != nil {
		wrappedErr := fmt.Errorf("error creating a new ConsumerGroup | err = %w", err)
		logger.Error().Err(wrappedErr).Send()
		panic(wrappedErr)
	}

	handlerWrapper := newConsumerGroupHandlerWrapper(processor)

	impl := consumerWrapperImpl{
		config:                      &config,
		consumerGroup:               consumerGroup,
		funcErrorHandling:           noopFuncError,
		consumerGroupHandlerWrapper: handlerWrapper,
		hasStarted:                  atomic.Bool{},
		hasStopped:                  atomic.Bool{},
		stopChan:                    make(chan struct{}),
		errorCount:                  atomic.Uint32{},
		topics:                      config.Common.Topics,
		principal:                   config.Principal,
		durationToResetCounter:      DefaultTimerResetTime,
		maxRestarts:                 atomic.Uint32{},
	}

	if config.Common.DurationToResetCounter != nil {
		impl.durationToResetCounter = *config.Common.DurationToResetCounter
	}
	impl.setMaxRestarts(config.Common)
	impl.hasStopped.Store(false)
	return &impl
}

func NewConsumerWrapperSaslSslAutoStart(config ConsumerGroupConfigSasl,
	processor ConsumedMessageProcessor) ConsumerWrapper {

	consumerWrapper := NewConsumerWrapperSaslSsl(config, processor)
	consumerWrapper.Start()
	return consumerWrapper
}
