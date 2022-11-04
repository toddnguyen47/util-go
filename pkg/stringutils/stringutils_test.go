package stringutils

import (
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
