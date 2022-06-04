package convertslicestomap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GivenStringSlice_WhenConverting_ThenConvertToMap(t *testing.T) {
	inputList := []string{"hello", "world", "42"}

	resultsMap := Convert(inputList)

	expectedMap := make(map[string]string)
	expectedMap["hello"] = "hello"
	expectedMap["world"] = "world"
	expectedMap["42"] = "42"
	assert.Equal(t, expectedMap, resultsMap)
}

func Test_GivenEmptyStringSlice_WhenConverting_ThenConvertToMap(t *testing.T) {
	var inputList []string

	resultsMap := Convert(inputList)

	expectedMap := make(map[string]string)
	assert.Equal(t, expectedMap, resultsMap)
}

func Test_GivenNilStringSlice_WhenConverting_ThenConvertToMap(t *testing.T) {
	resultsMap := Convert(nil)

	expectedMap := make(map[string]string)
	assert.Equal(t, expectedMap, resultsMap)
}
