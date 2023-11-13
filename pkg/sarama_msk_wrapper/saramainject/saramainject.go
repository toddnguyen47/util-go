package saramainject

import (
	"crypto/x509"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/toddnguyen47/util-go/pkg/loggerwrapper"
)

const (
	defaultBasePath           = "/tmp"
	defaultFileNameKerbKeytab = "ktab.keytab"
	defaultFileNameKerbConf   = "krb5.conf"
	defaultFileNameSslCert    = "ca-certificates.crt"

	_packageName = "saramainject"
)

// Monkey patching for tests
var (
	_osWriteFile = os.WriteFile
	_osMkdirAll  = os.MkdirAll
)

var (
	_wr            = diode.NewWriter(os.Stderr, 1000, 10*time.Millisecond, loggerwrapper.MissedLogger)
	_packageLogger = zerolog.New(_wr).With().Timestamp().Logger().Level(zerolog.DebugLevel)
)

type InjectPath struct {
	KerbKeytab string
	KerbConf   string
	SslCert    string
}

func TmpCertFolder(principal string) string {
	escapedPrincipal := url.QueryEscape(principal)
	s1 := fmt.Sprintf("%s/esc_%s", defaultBasePath, escapedPrincipal)
	return s1
}

// Inject - inject kerbKeytab, kerbConf, and sslCerts. Return the path in order as well.
func Inject(principal string, kerbKeytab []byte, kerbConf []byte, sslCerts []byte) InjectPath {
	logger := _packageLogger.With().Str("functionName", _packageName+":inject()").Logger()
	tmpCertFolder := TmpCertFolder(principal)
	logger.Info().Str("tmpCertFolder", tmpCertFolder).Msg("INJECTING!")
	const permission = 0666
	const formatStr = "%s/%s"

	err := _osMkdirAll(tmpCertFolder, 0750)
	if err != nil {
		logger.Error().Err(err).Msg("error making directory")
		panic(err)
	}

	// inject KERB_KEYTAB
	keyTabLocation := fmt.Sprintf(formatStr, tmpCertFolder, defaultFileNameKerbKeytab)
	err = _osWriteFile(keyTabLocation, kerbKeytab, 0666)
	if err != nil {
		newErr := fmt.Errorf("error wring KERB_KEYTAB; err: %w", err)
		logger.Error().Err(newErr).Send()
		panic(newErr)
	}

	// inject KERB_CONF
	var KerbConfLocation string
	if kerbConf != nil {
		KerbConfLocation = fmt.Sprintf(formatStr, tmpCertFolder, defaultFileNameKerbConf)
		err = _osWriteFile(KerbConfLocation, kerbConf, permission)
		if err != nil {
			newErr := fmt.Errorf("error writing KERB_CONF; err: %w", err)
			logger.Error().Err(newErr).Send()
			panic(newErr)
		}
	}

	// inject SSL_CERT
	var SslCertLocation string
	if sslCerts != nil {
		SslCertLocation = fmt.Sprintf(formatStr, tmpCertFolder, defaultFileNameSslCert)
		err = _osWriteFile(SslCertLocation, sslCerts, permission)
		if err != nil {
			newErr := fmt.Errorf("error writing SSL_CERT; err: %w", err)
			logger.Error().Err(newErr).Send()
			panic(newErr)
		}
	}

	return InjectPath{
		KerbKeytab: keyTabLocation,
		KerbConf:   KerbConfLocation,
		SslCert:    SslCertLocation,
	}
}

func GetCertsFrom(certsLocation string) (*x509.CertPool, error) {
	certPool := x509.NewCertPool()
	certs, err := os.ReadFile(certsLocation)
	if err != nil {
		return nil, err
	}
	if ok := certPool.AppendCertsFromPEM(certs); !ok {
		newErr := fmt.Errorf("cannot append certs to certPool")
		return nil, newErr
	}
	return certPool, nil
}
