package sarama_msk_wrapper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/toddnguyen47/util-go/pkg/loggerwrapper"
	"github.com/toddnguyen47/util-go/pkg/sarama_msk_wrapper/saramainject"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// /----------------------------------------------------------\
// #region public constants / variables
// ------------------------------------------------------------

const DefaultTimerResetTime = 30 * time.Minute

// ------------------------------------------------------------
// #endregion public constants / variables
// \----------------------------------------------------------/

var (
	_logLevel      = zerolog.WarnLevel
	_packageUuid   = uuid.New()
	_wr            = diode.NewWriter(os.Stderr, 1000, 10*time.Millisecond, loggerwrapper.MissedLogger)
	_packageLogger = zerolog.New(_wr).With().Timestamp().
			Str("packageUuid", _packageUuid.String()).Logger().Level(_logLevel)
	_terminationDelay = 5 * time.Second
	_printer          = message.NewPrinter(language.English)
	_certExtensions   = []string{".crt", ".keytab", ".conf"}
)

// Monkey patching for tests
var (
	_saramaNewAsyncProducer = sarama.NewAsyncProducer
	_saramaNewConsumerGroup = sarama.NewConsumerGroup
	_filepathGlob           = filepath.Glob
	_osRemove               = os.Remove
)

const (
	_packageNameAsyncProducer         = "sarama_msk_wrapper_asyncproducer"
	_packageNameAsyncProducerDisabled = "sarama_msk_wrapper_asyncproducer_disabled"
	_packageNameConsumerGroup         = "sarama_msk_wrapper_consumergroup"
	_packageNameConsumerGroupDisabled = "sarama_msk_wrapper_consumergroup_disabled"

	_formatDigit              = "%d"
	_strFunctionName          = "functionName"
	_defaultMaxRebalanceCount = uint32(10)
)

// SetLogLevel - Set LogLevel. Valid values can be found in zerolog.Level
// Reference: https://github.com/rs/zerolog
func SetLogLevel(level string) {
	//zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	lowercaseLevel := strings.ToLower(level)
	switch lowercaseLevel {
	case "debug":
		_logLevel = zerolog.DebugLevel
	case "info":
		_logLevel = zerolog.InfoLevel
	case "warn":
		_logLevel = zerolog.WarnLevel
	case "error":
		_logLevel = zerolog.ErrorLevel
	case "fatal":
		_logLevel = zerolog.FatalLevel
	case "panic":
		_logLevel = zerolog.PanicLevel
	case "nolevel":
		_logLevel = zerolog.NoLevel
	case "disabled":
		_logLevel = zerolog.Disabled
	case "trace":
		_logLevel = zerolog.TraceLevel
	}

	_packageLogger = _packageLogger.Level(_logLevel)
}

type configInterface interface {
	validate() error
	string() string
}

func deleteTmpCerts(principal string) {

	logger := getLoggerWithName("sarama_msk_wrapper:deleteTmpCerts()")
	filesList := make([]string, 0)
	tmpCertFolder := saramainject.TmpCertFolder(principal)
	for _, extension := range _certExtensions {
		files, err := _filepathGlob(tmpCertFolder + "/*" + extension)
		if err != nil {
			continue
		}
		filesList = append(filesList, files...)
	}
	for _, file := range filesList {
		logger.Info().Str("removing current file", file).Send()
		err := _osRemove(file)
		if err != nil {
			logger.Error().Err(err).Str("error removing file", file).Send()
			continue
		}
	}

	if _, err := os.Stat(tmpCertFolder); err == nil {
		// Folder exists
		logger.Info().Str("removing empty tmp folder", tmpCertFolder).Send()
		err = _osRemove(tmpCertFolder)
		if err != nil {
			newErr := fmt.Errorf("cannot remove tmpCertFolder; err: %w", err)
			logger.Error().Err(newErr).Send()
		}
	}
}

func getLoggerWithName(functionName string) zerolog.Logger {
	return _packageLogger.With().Str(_strFunctionName, functionName).Logger()
}

func noopFuncError(error) {}
