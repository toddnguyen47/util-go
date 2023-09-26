package saramaconfig

import (
	"crypto/tls"
	"fmt"
	"os"
	"strings"

	"github.com/IBM/sarama"
	"github.com/rs/zerolog"
	"github.com/toddnguyen47/util-go/pkg/sarama_msk_wrapper/saramainject"
)

// Monkey patching for tests
var (
	_getCertsFrom = saramainject.GetCertsFrom
)

var _packageLogger = zerolog.New(os.Stderr).With().Timestamp().Logger().Level(zerolog.InfoLevel)

// GetSaramaConfigSsl - return saramaConfig with SSL
func GetSaramaConfigSsl(pubKey []byte, privateKey []byte) *sarama.Config {

	saramaConfig := sarama.NewConfig()

	// Ref: https://github.com/Shopify/sarama/issues/2341
	// Load client cert
	cert, err := tls.X509KeyPair(pubKey, privateKey)
	if err != nil {
		newErr := fmt.Errorf("cannot load public and private key | err: %w", err)
		_packageLogger.Err(newErr).Send()
		panic(newErr)
	}

	// We want to use SSL TLS instead of SASL
	saramaConfig.Net.SASL.Enable = false
	saramaConfig.Net.TLS.Enable = true
	if saramaConfig.Net.TLS.Config == nil {
		saramaConfig.Net.TLS.Config = &tls.Config{}
	}
	saramaConfig.Net.TLS.Config.Certificates = append(saramaConfig.Net.TLS.Config.Certificates, cert)

	return saramaConfig
}

func GetSaramaConfigSasl(principal string, kerbKeytab []byte, kerbConf []byte, sslCert []byte) *sarama.Config {

	saramaConfig := sarama.NewConfig()

	injectPaths := saramainject.Inject(principal, kerbKeytab, kerbConf, sslCert)
	principalList := strings.Split(principal, "@")

	saramaConfig.Net.SASL.Enable = true
	saramaConfig.Net.SASL.Mechanism = sarama.SASLTypeGSSAPI
	saramaConfig.Net.SASL.GSSAPI.Realm = principalList[1]
	saramaConfig.Net.SASL.GSSAPI.Username = principalList[0]
	saramaConfig.Net.SASL.GSSAPI.KeyTabPath = injectPaths.KerbKeytab
	saramaConfig.Net.SASL.GSSAPI.ServiceName = "kafka"
	saramaConfig.Net.SASL.GSSAPI.AuthType = sarama.KRB5_KEYTAB_AUTH
	saramaConfig.Net.SASL.GSSAPI.DisablePAFXFAST = true

	if kerbConf != nil {
		saramaConfig.Net.SASL.GSSAPI.KerberosConfigPath = injectPaths.KerbConf
	}
	if sslCert != nil {
		certPool, certsErr := _getCertsFrom(injectPaths.SslCert)
		if certsErr != nil {
			panic(certsErr)
		}
		saramaConfig.Net.TLS.Enable = true
		saramaConfig.Net.TLS.Config = &tls.Config{
			ClientCAs:  certPool,
			MinVersion: tls.VersionTLS12,
		}
	}

	return saramaConfig
}
