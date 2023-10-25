package sarama_msk_wrapper

import (
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ConsumerGroupConfig struct {
	// Common - REQUIRED. Common configs for all consumer groups
	Common ConsumerGroupConfigCommon
	// PubKey - REQUIRED.
	PubKey []byte
	// PrivateKey - REQUIRED.
	PrivateKey []byte
}

func (c *ConsumerGroupConfig) validate() error {

	err := c.Common.validate()
	if err != nil {
		return err
	}

	return validation.ValidateStruct(c,
		validation.Field(&c.PubKey, validation.Required),
		validation.Field(&c.PrivateKey, validation.Required),
	)
}

func (c *ConsumerGroupConfig) validateBatch() error {
	err := c.validate()
	if err != nil {
		return err
	}
	return c.Common.validateBatch()
}

func (c *ConsumerGroupConfig) string() string {
	var sb strings.Builder
	sb.WriteString("brokers -> ")
	for i, elem := range c.Common.Brokers {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(elem)
	}
	sb.WriteString("; consumerGroupId -> " + c.Common.ConsumerGroupId)
	appendTopics(&sb, c.Common.Topics)
	return sb.String()
}

type ConsumerGroupConfigSasl struct {
	// Common - REQUIRED. Common configs for all consumer groups
	Common ConsumerGroupConfigCommon
	// Principal - REQUIRED. In the form of username@realm
	Principal string
	// KerbKeytab - REQUIRED. Base64 decoded byte.
	KerbKeytab []byte

	// Optionals

	// KerbConf - OPTIONAL. If no KerbConf is passed in, the default path from the platform team will be used.
	KerbConf []byte
	// SslCert - OPTIONAL. If no SslCert is passed in, the default path from the platform team will be used.
	SslCert []byte
}

func (c *ConsumerGroupConfigSasl) validate() error {
	err := c.Common.validate()
	if err != nil {
		return err
	}
	return validation.ValidateStruct(c,
		validation.Field(&c.Principal, validation.Required),
		validation.Field(&c.KerbKeytab, validation.Required),
	)
}

func (c *ConsumerGroupConfigSasl) string() string {
	var sb strings.Builder
	sb.WriteString("brokers -> ")
	for i, elem := range c.Common.Brokers {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(elem)
	}
	sb.WriteString("; principal -> " + c.Principal)
	sb.WriteString("; consumerGroupId -> " + c.Common.ConsumerGroupId)
	appendTopics(&sb, c.Common.Topics)
	return sb.String()
}

type ConsumerGroupConfigCommon struct {
	// Brokers - REQUIRED.
	Brokers []string
	// ConsumerGroupId - REQUIRED.
	ConsumerGroupId string
	// Topics - REQUIRED.
	Topics []string
	// BatchSize - REQUIRED only for batch
	BatchSize uint
	// BatchTimeout - REQUIRED only for batch
	BatchTimeout *time.Duration

	// DurationToResetCounter - OPTIONAL. Default to 30 minutes
	DurationToResetCounter *time.Duration
}

func (c *ConsumerGroupConfigCommon) validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Brokers, validation.Required),
		validation.Field(&c.ConsumerGroupId, validation.Required),
		validation.Field(&c.Topics, validation.Required),
	)
}

func (c *ConsumerGroupConfigCommon) validateBatch() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.BatchSize, validation.Required, validation.Min(uint(1))),
		validation.Field(&c.BatchTimeout, validation.Required),
	)
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
