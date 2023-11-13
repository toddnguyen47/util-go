package loggerwrapper

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog"
)

func MissedLogger(missed int) {
	fmt.Printf("Logger dropped %d messages\n", missed)
}

func (i1 *impl) GetLoggerWithName(functionName string) zerolog.Logger {
	return i1.packageLogger.With().Str("functionName", functionName).Logger()
}

// SetLogLevel - Set LogLevel. Valid values can be found in zerolog.Level
// Reference: https://github.com/rs/zerolog
func (i1 *impl) SetLogLevel(level string) {
	i1.logLevel = GetLogLevelFromString(level)
	i1.packageLogger = i1.packageLogger.Level(i1.logLevel)
}

func (i1 *impl) GetLogLevel() string {
	s1 := i1.logLevel.String()
	return strings.ToLower(s1)
}

func GetLogLevelFromString(level string) zerolog.Level {
	lowercaseLevel := strings.ToLower(level)
	logLevel := zerolog.WarnLevel
	switch lowercaseLevel {
	case "debug":
		logLevel = zerolog.DebugLevel
	case "info":
		logLevel = zerolog.InfoLevel
	case "warn":
		logLevel = zerolog.WarnLevel
	case "error":
		logLevel = zerolog.ErrorLevel
	case "fatal":
		logLevel = zerolog.FatalLevel
	case "panic":
		logLevel = zerolog.PanicLevel
	case "nolevel":
		logLevel = zerolog.NoLevel
	case "disabled":
		logLevel = zerolog.Disabled
	case "trace":
		logLevel = zerolog.TraceLevel
	}
	return logLevel
}
