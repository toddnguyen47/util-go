package multithreadcount

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GivenSetError_When_ThenGetErrorCorrectly(t *testing.T) {
	err := errors.New("some error")
	multiCount := NewMultiThreadCount()
	multiCount.SetError(err)

	assert.Equal(t, uint32(0), multiCount.LoadSuccessCount())
	assert.Equal(t, uint32(0), multiCount.LoadErrorCount())
	assert.Equal(t, err, multiCount.GetError())
}
