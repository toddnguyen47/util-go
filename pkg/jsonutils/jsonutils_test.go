package jsonutils

import (
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

func makePtr(_ *testing.T, str string) *string {
	return &str
}
