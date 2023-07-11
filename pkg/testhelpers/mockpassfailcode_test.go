package testhelpers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GivenFFP_ThenFailsTwiceThenPass(t *testing.T) {
	// -- ARRANGE --
	sutMockPassFail := NewMockPassFail()
	sutMockPassFail.SetCode("FFP")
	// -- ACT --
	// -- ASSERT --
	err := sutMockPassFail.WillPass()
	assert.True(t, errors.Is(err, ErrFunctionShouldFail), "err should be ErrFunctionShouldFail")
	err = sutMockPassFail.WillPass()
	assert.True(t, errors.Is(err, ErrFunctionShouldFail), "err should be ErrFunctionShouldFail")
	err = sutMockPassFail.WillPass()
	assert.Nil(t, err)
	err = sutMockPassFail.WillPass()
	assert.Nil(t, err)
	assert.Equal(t, 4, sutMockPassFail.GetCount())
}
