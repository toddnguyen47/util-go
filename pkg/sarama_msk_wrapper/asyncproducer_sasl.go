package sarama_msk_wrapper

import (
	"fmt"
	"sync/atomic"

	"github.com/toddnguyen47/util-go/pkg/sarama_msk_wrapper/saramaconfig"
)

func NewAsyncProducerSaslSsl(config AsyncProducerConfigSasl) AsyncProducerWrapper {

	logger := getLoggerWithName(_packageNameAsyncProducer + ":NewAsyncProducerSaslSsl()")
	err := config.validate()
	if err != nil {
		wrappedErr := fmt.Errorf("async producer config SASL_SSL error | err: %w", err)
		logger.Error().Err(wrappedErr).Send()
		panic(wrappedErr)
	}

	saramaConfig := saramaconfig.GetSaramaConfigSasl(config.Principal,
		config.KerbKeytab, config.KerbConf, config.SslCert)
	asyncProducer, err := _saramaNewAsyncProducer(config.Common.Brokers, saramaConfig)
	if err != nil {
		wrappedErr := fmt.Errorf("error creating a new AsyncProducer | err: %w", err)
		logger.Error().Err(wrappedErr).Send()
		panic(wrappedErr)
	}

	impl := asyncProducerImpl{
		config:                   &config,
		asyncProducer:            asyncProducer,
		stopChan:                 make(chan struct{}, 1),
		hasStopped:               atomic.Bool{},
		funcMetricErrorProducing: noopFunc,
		principal:                config.Principal,
		durationToResetCounter:   DefaultTimerResetTime,
	}
	impl.hasStopped.Store(false)
	if config.DurationToResetCounter != nil {
		impl.durationToResetCounter = *config.DurationToResetCounter
	}
	return &impl
}

func NewAsyncProducerSasSslAutoStart(config AsyncProducerConfigSasl) AsyncProducerWrapper {
	asyncProducerWrapper := NewAsyncProducerSaslSsl(config)
	asyncProducerWrapper.Start()
	return asyncProducerWrapper
}
