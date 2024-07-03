package testhelpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GivenErrReadCloser_ThenReturnErr(t *testing.T) {
	// -- GIVEN --
	closerErr := ErrReadCloser(1)
	// -- WHEN --
	n, err := closerErr.Read([]byte(""))
	// -- THEN --
	assert.NotNil(t, err)
	assert.Equal(t, 0, n)
	err = closerErr.Close()
	assert.Nil(t, err)
}
