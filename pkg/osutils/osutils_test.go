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

func Test_GivenNonExistentFile_ThenDoNotRemove(t *testing.T) {
	err := RemoveIfExists("asdf")
	assert.Nil(t, err)
}

func Test_GivenExistentFile_ThenRemove(t *testing.T) {
	filePath := "helloWorld.txt"
	err := os.WriteFile(filePath, []byte("hello world"), 0666)
	assert.Nil(t, err)
	err = RemoveIfExists(filePath)
	assert.Nil(t, err)
}
