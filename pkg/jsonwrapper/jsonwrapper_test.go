package jsonwrapper

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Id   *string `json:"id"`
	Name *string `json:"name"`
}

func Test_GivenValidMarshal_When_ThenErrIsNil(t *testing.T) {
	// -- ARRANGE --
	jsonWrapper := NewDefaultJsonWrapper()
	test1 := testStruct{
		Id:   makePtr(t, "id"),
		Name: nil,
	}
	// -- ACT --
	bytes1, err := jsonWrapper.Marshal(test1)
	// -- ASSERT --
	assert.Nil(t, err)
	assert.Greater(t, len(string(bytes1)), 0)
}

func Test_GivenInvalidMarshal_When_ThenErrIsNil(t *testing.T) {
	// -- ARRANGE --
	jsonWrapper := NewDefaultJsonWrapper()
	chan1 := make(chan string)
	// -- ACT --
	bytes1, err := jsonWrapper.Marshal(chan1)
	// -- ASSERT --
	assert.NotNil(t, err)
	assert.Equal(t, len(string(bytes1)), 0)
}

func Test_GivenValidUnmarshal_When_ThenErrIsNil(t *testing.T) {
	// -- ARRANGE --
	jsonWrapper := NewDefaultJsonWrapper()
	test1 := testStruct{
		Id:   makePtr(t, "id"),
		Name: nil,
	}
	bytes1, err := jsonWrapper.Marshal(test1)
	assert.Nil(t, err)
	// -- ACT --
	var output map[string]interface{}
	err = jsonWrapper.Unmarshal(bytes1, &output)
	// -- ASSERT --
	assert.Nil(t, err)
}

func Test_GivenInvalidUnmarshal_When_ThenErrIsNil(t *testing.T) {
	// -- ARRANGE --
	jsonWrapper := NewDefaultJsonWrapper()
	test1 := testStruct{
		Id:   makePtr(t, "id"),
		Name: nil,
	}
	bytes1, err := jsonWrapper.Marshal(test1)
	assert.Nil(t, err)
	// -- ACT --
	chan1 := make(chan string)
	err = jsonWrapper.Unmarshal(bytes1, &chan1)
	// -- ASSERT --
	assert.NotNil(t, err)
}

func Test_GivenEncodingProperly_ThenErrIsNil(t *testing.T) {
	// -- ARRANGE --
	a := testStruct{
		Id:   makePtr(t, "id"),
		Name: makePtr(t, "name & address"),
	}
	// -- ACT --
	b1, err := MarshalNoEscapeHtml(a)
	// -- ASSERT --
	assert.Nil(t, err)
	str1 := string(b1)
	assert.True(t, strings.Contains(str1, "name & address"), "should contain `&`")
}

func Test_GivenEncodingImproperly_ThenErrIsNotNil(t *testing.T) {
	// -- ARRANGE --
	a := make(chan string)
	// -- ACT --
	b1, err := MarshalNoEscapeHtml(a)
	// -- ASSERT --
	assert.NotNil(t, err)
	assert.Equal(t, []byte{}, b1)
}

func makePtr(t *testing.T, str string) *string {
	return &str
}
