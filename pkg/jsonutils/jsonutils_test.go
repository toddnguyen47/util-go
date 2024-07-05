package jsonutils

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Id   *string `json:"id"`
	Name *string `json:"name"`
}

func Test_GivenEncodingProperly_ThenErrIsNil(t *testing.T) {
	// -- GIVEN --
	a := testStruct{
		Id:   makePtr(t, "id"),
		Name: makePtr(t, "name & address"),
	}
	// -- WHEN --
	b1, err := MarshalNoEscapeHtml(a)
	// -- THEN --
	assert.Nil(t, err)
	str1 := string(b1)
	assert.True(t, strings.Contains(str1, "name & address"), "should contain `&`")
}

func Test_GivenEncodingImproperly_ThenErrIsNotNil(t *testing.T) {
	// -- GIVEN --
	a := make(chan string)
	// -- WHEN --
	b1, err := MarshalNoEscapeHtml(a)
	// -- THEN --
	assert.NotNil(t, err)
	assert.Equal(t, []byte{}, b1)
}

func Test_GivenSimpleJsonData_ThenIterateProperly(t *testing.T) {
	// -- GIVEN --
	data := []byte(`{"key": [1,2,3], "key2": null}`)
	var inputData map[string]any
	err := json.Unmarshal(data, &inputData)
	assert.NoError(t, err)
	// -- WHEN --
	map1 := IterateJson(inputData)
	// -- THEN --
	assert.Equal(t, map1["$.key[0]"], float64(1))
	assert.Equal(t, map1["$.key[1]"], float64(2))
	assert.Equal(t, map1["$.key[2]"], float64(3))
}

func Test_GivenComplexJsonData_ThenIterateProperly(t *testing.T) {
	// -- GIVEN --
	data, err := os.ReadFile("testdata/example1.json")
	assert.NoError(t, err)
	var inputData map[string]any
	err = json.Unmarshal(data, &inputData)
	assert.NoError(t, err)
	// -- WHEN --
	map1 := IterateJson(inputData)
	// -- THEN --
	assert.Equal(t, map1["$.web-app.servlet[0].servlet-class"], "org.cofax.cds.CDSServlet")
	assert.Equal(t, map1["$.web-app.servlet[4].init-param.omegaServer"], nil)
}

func makePtr(_ *testing.T, str string) *string {
	return &str
}
