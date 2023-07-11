package testhelpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GivenErrReadCloser_ThenReturnErr(t *testing.T) {
	// -- ARRANGE --
	closerErr := ErrReadCloser(1)
	// -- ACT --
	n, err := closerErr.Read([]byte(""))
	// -- ASSERT --
	assert.NotNil(t, err)
	assert.Equal(t, 0, n)
	err = closerErr.Close()
	assert.Nil(t, err)
}
