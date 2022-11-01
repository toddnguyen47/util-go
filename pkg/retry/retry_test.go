package retry

import (
	"errors"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func testFunction(input int) (int, error) {
	if input < 42 {
		return input, errors.New("some error")
	}
	return input, nil
}

func wrapper(arguments ...interface{}) error {
	_, err := testFunction(arguments[0].(int))
	return err
}

func Test_GivenValidTestFunction_When_ThenRetryA5Times(t *testing.T) {
	input := 45
	logger := zerolog.Logger{}

	err := IncrementalRetry(&logger, 5, 100*time.Millisecond, wrapper, input)

	assert.Nil(t, err)
}

func Test_GivenValidTestFunctionExceedMaxRetryTimes_When_ThenRetryA5Times(t *testing.T) {
	input := 30
	logger := zerolog.Logger{}

	err := IncrementalRetry(&logger, 5, 100*time.Millisecond, wrapper, input)

	assert.NotNil(t, err)
}
