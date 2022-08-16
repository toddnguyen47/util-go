package stringutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GivenNilPtr_When_ThenIsBlankIsTrue(t *testing.T) {
	assert.True(t, IsBlank(nil))
}

func Test_GivenEmptyString_When_ThenIsBlankIsTrue(t *testing.T) {
	a := ""
	assert.True(t, IsBlank(&a))
}

func Test_GivenOnlyWhitespace_When_ThenIsBlankIsTrue(t *testing.T) {
	a := " \t "
	assert.True(t, IsBlank(&a))
}

func Test_GivenNonWhitespace_When_ThenIsBlankIsFalse(t *testing.T) {
	a := "Bob"
	assert.False(t, IsBlank(&a))
}

func Test_GivenNonWhitespaceWithSomeWhitespace_When_ThenIsBlankIsFalse(t *testing.T) {
	a := "   Bob  \t  "
	assert.False(t, IsBlank(&a))
}
