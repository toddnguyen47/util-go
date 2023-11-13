package loggerwrapper

import (
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/toddnguyen47/util-go/pkg/sarama_msk_wrapper/saramainject"
)

type Interface interface {
	GetLoggerWithName(functionName string) zerolog.Logger

	GetLogLevel() string
	SetLogLevel(level string)
}

type impl struct {
	logLevel      zerolog.Level
	packageLogger zerolog.Logger
}

func NewLoggerWrapper() Interface {
	logLevel := zerolog.WarnLevel
	packageUuid := uuid.New()
	wr := diode.NewWriter(os.Stderr, 1000, 10*time.Millisecond, saramainject.MissedLogger)
	packageLogger := zerolog.New(wr).With().Timestamp().
		Str("packageUuid", packageUuid.String()).Logger().Level(logLevel)

	i1 := impl{
		logLevel:      logLevel,
		packageLogger: packageLogger,
	}
	return &i1
}
