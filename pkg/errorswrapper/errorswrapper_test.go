package errorswrapper

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GivenErrorsWrapper_ThenWrapProperly(t *testing.T) {
	err1 := errors.New("err1")
	err2 := Wrap(err1, "err2")
	assert.True(t, errors.Is(err2, err1))

	err3 := Wrap(err2, "err3")
	assert.True(t, errors.Is(err3, err2))
}

func Test_GivenErrNil_ThenReturnProperError(t *testing.T) {
	var err error
	err2 := Wrap(err, "new message")
	assert.Equal(t, "new message", err2.Error())
}
