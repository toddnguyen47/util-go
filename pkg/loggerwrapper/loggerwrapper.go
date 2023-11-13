package loggerwrapper

import (
	"strings"

	"github.com/rs/zerolog"
)

func (i1 *impl) GetLoggerWithName(functionName string) zerolog.Logger {
	return i1.packageLogger.With().Str("functionName", functionName).Logger()
}

// SetLogLevel - Set LogLevel. Valid values can be found in zerolog.Level
// Reference: https://github.com/rs/zerolog
func (i1 *impl) SetLogLevel(level string) {
	//zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	lowercaseLevel := strings.ToLower(level)
	switch lowercaseLevel {
	case "debug":
		i1.logLevel = zerolog.DebugLevel
	case "info":
		i1.logLevel = zerolog.InfoLevel
	case "warn":
		i1.logLevel = zerolog.WarnLevel
	case "error":
		i1.logLevel = zerolog.ErrorLevel
	case "fatal":
		i1.logLevel = zerolog.FatalLevel
	case "panic":
		i1.logLevel = zerolog.PanicLevel
	case "nolevel":
		i1.logLevel = zerolog.NoLevel
	case "disabled":
		i1.logLevel = zerolog.Disabled
	case "trace":
		i1.logLevel = zerolog.TraceLevel
	}

	i1.packageLogger = i1.packageLogger.Level(i1.logLevel)
}

func (i1 *impl) GetLogLevel() string {
	s1 := i1.logLevel.String()
	return strings.ToLower(s1)
}
