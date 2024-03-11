package stringutils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GivenNilPtr_When_ThenIsBlankIsTrue(t *testing.T) {
	assert.True(t, IsBlank(nil))
	assert.False(t, IsNotBlank(nil))
}

func Test_GivenEmptyString_When_ThenIsBlankIsTrue(t *testing.T) {
	a := ""
	assert.True(t, IsBlank(&a))
	assert.False(t, IsNotBlank(&a))
}

func Test_GivenOnlyWhitespace_When_ThenIsBlankIsTrue(t *testing.T) {
	a := " \t "
	assert.True(t, IsBlank(&a))
	assert.False(t, IsNotBlank(&a))
}

func Test_GivenNonWhitespace_When_ThenIsBlankIsFalse(t *testing.T) {
	a := "Bob"
	assert.False(t, IsBlank(&a))
	assert.True(t, IsNotBlank(&a))
}

func Test_GivenNonWhitespaceWithSomeWhitespace_When_ThenIsBlankIsFalse(t *testing.T) {
	a := "   Bob  \t  "
	assert.False(t, IsBlank(&a))
	assert.True(t, IsNotBlank(&a))
}

func Test_GivenStr_WhenMakingPtr_ThenReturnStringPtr(t *testing.T) {
	a := "Bob"
	assert.Equal(t, &a, MakePtr(a))
}

func Test_GivenElems2Empties_WhenJoinExcludeEmpty_ThenReturnProperString(t *testing.T) {
	// -- GIVEN --
	strings := []string{"", "hello", "    ", "   world  ", "       "}
	// -- WHEN --
	str1 := JoinExcludeEmpty(strings, ";")
	// -- THEN --
	assert.Equal(t, "hello;   world  ", str1)
}

func Test_GivenStrGreaterThanIndex_WhenGetSubstring_ThenReturnProperSubstr(t *testing.T) {
	// -- GIVEN --
	strInput := "Doloremque eligendi est aut aut sint animi vitae voluptates"
	// -- WHEN --
	substr := GetSubstring(strInput, len(strInput)+500)
	// -- THEN --
	assert.Equal(t, strInput, substr)
}

func Test_GivenStrEqualToIndex_WhenGetSubstring_ThenReturnProperSubstr(t *testing.T) {
	// -- GIVEN --
	strInput := "Doloremque eligendi est aut aut sint animi vitae voluptates"
	// -- WHEN --
	substr := GetSubstring(strInput, len(strInput))
	// -- THEN --
	assert.Equal(t, strInput, substr)
}

func Test_GivenStrSmallerThanToIndex_WhenGetSubstring_ThenReturnProperSubstr(t *testing.T) {
	// -- GIVEN --
	strInput := "Doloremque eligendi est aut aut sint animi vitae voluptates"
	// -- WHEN --
	substr := GetSubstring(strInput, 11)
	// -- THEN --
	assert.Equal(t, "Doloremque ", substr)
}

func Test_GiveIndexIsNegative_WhenGetSubstring_ThenReturnProperSubstr(t *testing.T) {
	// -- GIVEN --
	strInput := "Doloremque eligendi est aut aut sint animi vitae voluptates"
	// -- WHEN --
	substr := GetSubstring(strInput, -1)
	// -- THEN --
	assert.Equal(t, strInput, substr)
}

func Test_GivenFirstWrite_WhenWriteToSbWithSep_ThenDoNotWriteSep(t *testing.T) {
	// -- GIVEN --
	sep := ","
	strInput := "hello"
	var sb strings.Builder
	// -- WHEN --
	WriteToSbWithSep(&sb, strInput, sep)
	// -- THEN --
	assert.False(t, strings.Contains(sb.String(), sep))
}

func Test_GivenSecondWrite_WhenWriteToSbWithSep_ThenDoWriteSep(t *testing.T) {
	// -- GIVEN --
	sep := ","
	strInput := "hello"
	var sb strings.Builder
	// -- WHEN --
	WriteToSbWithSep(&sb, strInput, sep)
	WriteToSbWithSep(&sb, strInput, sep)
	// -- THEN --
	assert.True(t, strings.Contains(sb.String(), sep))
}
