package sarama_msk_wrapper

import (
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AsyncProducerConfig struct {
	// Required

	// Brokers - REQUIRED.
	Brokers []string
	// PubKey - REQUIRED.
	PubKey []byte
	// PrivateKey - REQUIRED.
	PrivateKey []byte

	// Optionals

	// DurationToResetCounter - OPTIONAL. Default to 30 minutes
	DurationToResetCounter *time.Duration
}

func (c *AsyncProducerConfig) validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Brokers, validation.Required),
		validation.Field(&c.PubKey, validation.Required),
		validation.Field(&c.PrivateKey, validation.Required),
	)
}

func (c *AsyncProducerConfig) string() string {
	var sb strings.Builder
	sb.WriteString("brokers -> ")
	for i, elem := range c.Brokers {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(elem)
	}
	return sb.String()
}

type AsyncProducerConfigSasl struct {

	// Required

	// Principal - REQUIRED. In the form of username@realm
	Principal string
	// Brokers - REQUIRED.
	Brokers []string
	// KerbKeytab - REQUIRED. Base64 decoded byte.
	KerbKeytab []byte

	// Optionals

	// KerbConf - OPTIONAL. If no KerbConf is passed in, the default path from the platform team will be used.
	KerbConf []byte
	// SslCert - OPTIONAL. If no SslCert is passed in, the default path from the platform team will be used.
	SslCert []byte
	// DurationToResetCounter - OPTIONAL. Default to 30 minutes
	DurationToResetCounter *time.Duration
}

func (c *AsyncProducerConfigSasl) validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Principal, validation.Required),
		validation.Field(&c.Brokers, validation.Required),
		validation.Field(&c.KerbKeytab, validation.Required),
	)
}

func (c *AsyncProducerConfigSasl) string() string {
	var sb strings.Builder
	sb.WriteString("brokers -> ")
	for i, elem := range c.Brokers {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(elem)
	}
	sb.WriteString("; principal -> " + c.Principal)
	return sb.String()
}
