package loggerwrapper

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// /--------------------------------------------------------------------------\
// #region SETUP
// ----------------------------------------------------------------------------

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type LoggerWrapperTestSuite struct {
	suite.Suite
	ctxBg context.Context
}

func (s *LoggerWrapperTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
}

func (s *LoggerWrapperTestSuite) TearDownTest() {
	s.resetMonkeyPatching()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestLoggerWrapperTestSuite(t *testing.T) {
	suite.Run(t, new(LoggerWrapperTestSuite))
}

// ----------------------------------------------------------------------------
// #endregion SETUP
// \--------------------------------------------------------------------------/

// /--------------------------------------------------------------------------\
// #region TESTS ARE BELOW
// ----------------------------------------------------------------------------

func (s *LoggerWrapperTestSuite) Test_GivenLevel_ThenGetLogLevelProperly() {
	// -- ARRANGE --
	loggerWrapper := NewLoggerWrapper()
	// -- ACT --
	// -- ASSERT --
	logger := loggerWrapper.GetLoggerWithName("helloWorld")
	assert.NotNil(s.T(), logger)
	loggerWrapper2 := NewLoggerWrapperLogLevel("warn")
	assert.NotNil(s.T(), loggerWrapper2.GetLogLevel())
	// just to test MissedLogger
	MissedLogger(5)
	// -- ASSERT ALL LOG LEVELS --
	levels := []string{"debug", "info", "warn", "error", "fatal", "panic", "nolevel", "disabled", "trace"}
	for i := 0; i < len(levels); i++ {
		curLevel := levels[i]
		loggerWrapper.SetLogLevel(curLevel)
		if strings.EqualFold("nolevel", curLevel) {
			assert.Equal(s.T(), "", loggerWrapper.GetLogLevel())
		} else {
			assert.Equal(s.T(), curLevel, loggerWrapper.GetLogLevel())
		}
	}
}

// ----------------------------------------------------------------------------
// #endregion TESTS
// \--------------------------------------------------------------------------/

// /--------------------------------------------------------------------------\
// #region TEST HELPERS
// ----------------------------------------------------------------------------

func (s *LoggerWrapperTestSuite) resetMonkeyPatching() {
}

// ----------------------------------------------------------------------------
// #region TEST HELPERS
// \--------------------------------------------------------------------------/
