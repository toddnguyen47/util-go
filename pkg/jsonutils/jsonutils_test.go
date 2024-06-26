package jsonutils

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Id   *string `json:"id"`
	Name *string `json:"name"`
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

func Test_GivenSimpleJsonData_ThenIterateProperly(t *testing.T) {
	// -- GIVEN --
	data := []byte(`{"key": [1,2,3]}`)
	var inputData map[string]any
	err := json.Unmarshal(data, &inputData)
	assert.NoError(t, err)
	map1 := make(map[string]any)
	// -- WHEN --
	IterateJson(inputData, []string{}, map1)
	// -- THEN --
	assert.Equal(t, map1["key.0"], float64(1))
	assert.Equal(t, map1["key.1"], float64(2))
	assert.Equal(t, map1["key.2"], float64(3))
}

func makePtr(_ *testing.T, str string) *string {
	return &str
}
