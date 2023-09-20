package sarama_msk_wrapper

import (
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ConsumerGroupConfig struct {
	// Required

	// Brokers - REQUIRED.
	Brokers []string
	// PubKey - REQUIRED.
	PubKey []byte
	// PrivateKey - REQUIRED.
	PrivateKey []byte
	// ConsumerGroupId - REQUIRED.
	ConsumerGroupId string
	// Topics - REQUIRED.
	Topics []string
	// BatchSize - REQUIRED only for batch
	BatchSize uint
	// BatchTimeout - REQUIRED only for batch
	BatchTimeout *time.Duration

	// Optionals

	// DurationToResetCounter - OPTIONAL. Default to 30 minutes
	DurationToResetCounter *time.Duration
}

func (c *ConsumerGroupConfig) validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Brokers, validation.Required),
		validation.Field(&c.PubKey, validation.Required),
		validation.Field(&c.PrivateKey, validation.Required),
		validation.Field(&c.ConsumerGroupId, validation.Required),
		validation.Field(&c.Topics, validation.Required),
	)
}

func (c *ConsumerGroupConfig) validateBatch() error {
	err := c.validate()
	if err != nil {
		return err
	}
	return validation.ValidateStruct(c,
		validation.Field(&c.BatchSize, validation.Required, validation.Min(uint(0))),
		validation.Field(&c.BatchTimeout, validation.Required),
	)
}

func (c *ConsumerGroupConfig) string() string {
	var sb strings.Builder
	sb.WriteString("brokers -> ")
	for i, elem := range c.Brokers {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(elem)
	}
	sb.WriteString("; consumerGroupId -> " + c.ConsumerGroupId)
	appendTopics(&sb, c.Topics)
	return sb.String()
}

type ConsumerGroupConfigSasl struct {
	// Required

	// Principal - REQUIRED. In the form of username@realm
	Principal string
	// Brokers - REQUIRED.
	Brokers []string
	// KerbKeytab - REQUIRED. Base64 decoded byte.
	KerbKeytab []byte
	// ConsumerGroupId - REQUIRED.
	ConsumerGroupId string
	// Topics - REQUIRED.
	Topics []string

	// Optionals

	// KerbConf - OPTIONAL. If no KerbConf is passed in, the default path from the platform team will be used.
	KerbConf []byte
	// SslCert - OPTIONAL. If no SslCert is passed in, the default path from the platform team will be used.
	SslCert []byte
	// DurationToResetCounter - OPTIONAL. Default to 30 minutes
	DurationToResetCounter *time.Duration
}

func (c *ConsumerGroupConfigSasl) validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Principal, validation.Required),
		validation.Field(&c.Brokers, validation.Required),
		validation.Field(&c.KerbKeytab, validation.Required),
		validation.Field(&c.ConsumerGroupId, validation.Required),
		validation.Field(&c.Topics, validation.Required),
	)
}

func (c *ConsumerGroupConfigSasl) string() string {
	var sb strings.Builder
	sb.WriteString("brokers -> ")
	for i, elem := range c.Brokers {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(elem)
	}
	sb.WriteString("; principal -> " + c.Principal)
	sb.WriteString("; consumerGroupId -> " + c.ConsumerGroupId)
	appendTopics(&sb, c.Topics)
	return sb.String()
}

func appendTopics(sb *strings.Builder, topics []string) {
	sb.WriteString("; topics -> ")
	for i, elem := range topics {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(elem)
	}
}
