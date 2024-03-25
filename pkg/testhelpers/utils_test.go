package testhelpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GivenNoFile_ThenPanic(t *testing.T) {
	assert.Panics(t, func() {
		ReadInTestFile("non-existent-file")
	})
}

func Test_GivenFile_ThenNoPanic(t *testing.T) {
	b1 := ReadInTestFile("LICENSE")
	assert.NotEqual(t, []byte{}, b1)
}
