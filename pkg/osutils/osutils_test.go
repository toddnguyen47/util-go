package osutils

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	testKey      = "testKey"
	testValue    = "testValue"
	defaultValue = "defaultValue"
)

// ############################################################################
// #region SETUP
// ############################################################################

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type OsUtilsTestSuite struct {
	suite.Suite
	ctxBg context.Context
}

func (s *OsUtilsTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
	err := os.Setenv(testKey, "")
	assert.Nil(s.T(), err)
}

func (s *OsUtilsTestSuite) TearDownTest() {
	s.resetMonkeyPatching()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestOsUtilsTestSuite(t *testing.T) {
	suite.Run(t, new(OsUtilsTestSuite))
}

// #endregion

// ############################################################################
// #region TESTS ARE BELOW
// ############################################################################

func (s *OsUtilsTestSuite) Test_GivenEnvVarExists_ThenReturnVal() {
	err := os.Setenv(testKey, testValue)
	assert.Nil(s.T(), err)
	val := GetEnvWithDefault(testKey, defaultValue)
	assert.Equal(s.T(), testValue, val)
}

func (s *OsUtilsTestSuite) Test_GivenEnvVarDoesNotExist_ThenReturnDefault() {
	val := GetEnvWithDefault(testKey, defaultValue)
	assert.Equal(s.T(), defaultValue, val)
}

func (s *OsUtilsTestSuite) Test_GivenNonExistentFile_ThenDoNotRemove() {
	err := RemoveIfExists("asdf")
	assert.Nil(s.T(), err)
}

func (s *OsUtilsTestSuite) Test_GivenEnvVarExists_WhenInt_ThenReturnVal() {
	err := os.Setenv(testKey, "5")
	assert.Nil(s.T(), err)
	val := GetEnvWithDefaultInt(testKey, 100)
	assert.Equal(s.T(), 5, val)
}

func (s *OsUtilsTestSuite) Test_GivenEnvVarDoesNotExists_WhenInt_ThenReturnDefaultValue() {
	val := GetEnvWithDefaultInt(testKey, 100)
	assert.Equal(s.T(), 100, val)
}

func (s *OsUtilsTestSuite) Test_GivenEnvVarExists_WhenFloat64_ThenReturnVal() {
	err := os.Setenv(testKey, "5.5")
	assert.Nil(s.T(), err)
	val := GetEnvWithDefaultFloat64(testKey, 42.2)
	assert.Equal(s.T(), 5.5, val)
}

func (s *OsUtilsTestSuite) Test_GivenEnvVarDoesNotExists_WhenFloat64_ThenReturnDefaultValue() {
	val := GetEnvWithDefaultFloat64(testKey, 42.2)
	assert.Equal(s.T(), 42.2, val)
}

func (s *OsUtilsTestSuite) Test_GivenEnvVarExists_WhenBool_ThenReturnVal() {
	err := os.Setenv(testKey, "true")
	assert.Nil(s.T(), err)
	val := GetEnvWithDefaultBool(testKey, false)
	assert.True(s.T(), val)
}

func (s *OsUtilsTestSuite) Test_GivenEnvVarDoesNotExists_WhenBool_ThenReturnDefaultValue() {
	val := GetEnvWithDefaultBool(testKey, false)
	assert.False(s.T(), val)
}

func (s *OsUtilsTestSuite) Test_GivenExistentFile_ThenRemove() {
	filePath := "helloWorld.txt"
	err := os.WriteFile(filePath, []byte("hello world"), 0666)
	assert.Nil(s.T(), err)
	err = RemoveIfExists(filePath)
	assert.Nil(s.T(), err)
}

// ############################################################################
// #region TEST HELPERS
// ############################################################################

func (s *OsUtilsTestSuite) resetMonkeyPatching() {
}

// #endregion
