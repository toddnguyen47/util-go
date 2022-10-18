package unmarshaltomap

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toddnguyen47/util-go/pkg/jsonwrapper"
)

type TestStruct struct {
	Id   *string `json:"id"`
	Name *string `json:"name"`
}

func Test_GivenValid_When_ThenReturnMap(t *testing.T) {
	// -- ARRANGE --
	testStruct := TestStruct{
		Id:   makePtr("id"),
		Name: nil,
	}
	jsonWrapper := jsonwrapper.NewDefaultJsonWrapper()
	// -- ACT --
	map1, err := UnmarshalToMap(testStruct, jsonWrapper)
	// -- ASSERT --
	assert.Nil(t, err)
	assert.Equal(t, *testStruct.Id, map1["id"])
	assert.Nil(t, map1["name"])
}

func Test_GivenJsonMarshalError_When_ThenReturnErr(t *testing.T) {
	// -- ARRANGE --
	testChan := make(chan string)
	errJson := new(ErrorJson)
	errJson.marshalError = true
	// -- ACT --
	map1, err := UnmarshalToMap(testChan, errJson)
	// -- ASSERT --
	assert.NotNil(t, err)
	assert.Equal(t, make(map[string]interface{}), map1)
}

func Test_GivenJsonUnmarshalError_When_ThenReturnErr(t *testing.T) {
	// -- ARRANGE --
	testStruct := TestStruct{
		Id:   makePtr("id"),
		Name: nil,
	}
	errJson := new(ErrorJson)
	errJson.unmarshalError = true
	// -- ACT --
	map1, err := UnmarshalToMap(testStruct, errJson)
	// -- ASSERT --
	assert.NotNil(t, err)
	assert.Equal(t, make(map[string]interface{}), map1)
}

func makePtr(str string) *string {
	return &str
}

type ErrorJson struct {
	jsonwrapper.Interface
	marshalError   bool
	unmarshalError bool
}

func (e *ErrorJson) Marshal(v interface{}) ([]byte, error) {
	if e.marshalError {
		return nil, errors.New("some error")
	}
	return json.Marshal(v)
}

func (e *ErrorJson) Unmarshal(data []byte, v interface{}) error {
	if e.unmarshalError {
		return errors.New("some error")
	}
	return json.Unmarshal(data, v)
}
