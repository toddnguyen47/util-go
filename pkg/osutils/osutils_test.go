package osutils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testKey = "testKey"
const testValue = "testValue"
const defaultValue = "defaultValue"

func SetUp(t *testing.T) {
	err := os.Setenv(testKey, "")
	assert.Nil(t, err)
}

func Test_GivenEnvVarExists_ThenReturnVal(t *testing.T) {
	SetUp(t)
	err := os.Setenv(testKey, testValue)
	assert.Nil(t, err)
	val := GetEnvWithDefault(testKey, defaultValue)
	assert.Equal(t, testValue, val)
}

func Test_GivenEnvVarDoesNotExist_ThenReturnDefault(t *testing.T) {
	SetUp(t)
	val := GetEnvWithDefault(testKey, defaultValue)
	assert.Equal(t, defaultValue, val)
}
