package multithreadcount

import (
	"errors"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GivenSetError_When_ThenGetErrorCorrectly(t *testing.T) {
	err := errors.New("some error")
	multiCount := NewMultiThreadCount()
	multiCount.SetError(err)

	assert.Equal(t, uint32(0), atomic.LoadUint32(multiCount.ErrCount))
	assert.Equal(t, uint32(0), atomic.LoadUint32(multiCount.SuccessCount))
	assert.Equal(t, err, multiCount.GetError())
}
