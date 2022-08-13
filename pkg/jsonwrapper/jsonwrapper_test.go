package jsonwrapper

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GivenJsonMarshal_When_ThenMarshalSuccessfully(t *testing.T) {
	data1 := []byte("{}")
	jsonWrapper := NewJsonWrapper()

	actualResults, actualErr := jsonWrapper.Marshal(data1)

	expectedResults, expectedErr := json.Marshal(data1)
	assert.Equal(t, expectedResults, actualResults)
	assert.Equal(t, expectedErr, actualErr)
}

func Test_GivenJsonUnmarshal_When_ThenUnmarshalSuccessfully(t *testing.T) {
	data1 := []byte("{}")
	jsonWrapper := NewJsonWrapper()
	var output interface{}
	var output2 interface{}

	actualErr := jsonWrapper.Unmarshal(data1, &output)

	expectedErr := json.Unmarshal(data1, &output2)
	assert.Equal(t, expectedErr, actualErr)
}
