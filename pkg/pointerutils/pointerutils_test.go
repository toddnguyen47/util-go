package pointerutils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GivenString_ThenReturnPointerString(t *testing.T) {
	input := "hello"
	assert.Equal(t, &input, PtrString(input))
}

func Test_GivenInt_ThenReturnPointerInt(t *testing.T) {
	input := 2
	assert.Equal(t, &input, PtrInt(input))
}

func Test_GivenInt32_ThenReturnPointerInt(t *testing.T) {
	input := int32(2)
	assert.Equal(t, &input, PtrInt32(input))
}

func Test_GivenInt64_ThenReturnPointerInt(t *testing.T) {
	input := int64(2)
	assert.Equal(t, &input, PtrInt64(input))
}

func Test_GivenBool_ThenReturnPointerBool(t *testing.T) {
	input := false
	assert.Equal(t, &input, PtrBool(input))
}

func Test_GivenDuration_ThenReturnPointerDuration(t *testing.T) {
	input := 5 * time.Millisecond
	assert.Equal(t, &input, PtrDuration(input))
}
