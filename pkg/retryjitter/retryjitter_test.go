package retryjitter

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errForTests = errors.New("errForTests")

type mockRetry struct {
	// "F" for F, "P" for pass
	stringCode string
}

func (m *mockRetry) myFunction() error {
	if len(m.stringCode) > 0 {
		firstChar := m.stringCode[0]
		m.stringCode = m.stringCode[1:]
		if firstChar == 'F' {
			return errForTests
		}
	}
	return nil
}

func Test_GivenRetrySuccess_ThenReturnNil(t *testing.T) {
	// -- GIVEN --
	resetMonkeyPatching(t)
	mr := new(mockRetry)
	mr.stringCode = "FFP"
	retryTimes := 3
	// -- WHEN --
	err := Retry(retryTimes, mr.myFunction)
	// -- THEN --
	assert.Nil(t, err)
}

func Test_GivenRetrySuccessButTimeoutLessThanZero_ThenTimeoutDefaultsTo100ReturnErrorNil(t *testing.T) {
	// -- GIVEN --
	resetMonkeyPatching(t)
	mr := new(mockRetry)
	mr.stringCode = "FFP"
	retryTimes := 3
	// -- WHEN --
	err := RetryWithTimeout(retryTimes, -500, mr.myFunction)
	// -- THEN --
	assert.Nil(t, err)
}

func Test_GivenRetrySuccessButMinSleepTimeGreaterThanMaxSleepTime_ThenTimeoutDefaultsTo100ReturnErrorNil(t *testing.T) {
	// -- GIVEN --
	resetMonkeyPatching(t)
	mr := new(mockRetry)
	mr.stringCode = "FFP"
	retryTimes := 3
	_minSleepTimeMillis = 50
	// -- WHEN --
	err := RetryWithTimeout(retryTimes, 1, mr.myFunction)
	// -- THEN --
	assert.Nil(t, err)
}

func Test_GivenGeneratingRandomIntErrorRetrySuccess_ThenReturnNil(t *testing.T) {
	// -- GIVEN --
	mr := new(mockRetry)
	mr.stringCode = "FFP"
	retryTimes := 3
	// -- WHEN --
	err := Retry(retryTimes, mr.myFunction)
	// -- THEN --
	assert.Nil(t, err)
}

func Test_GivenRetryFailure_ThenReturnErr(t *testing.T) {
	resetMonkeyPatching(t)
	// -- GIVEN --
	mr := new(mockRetry)
	mr.stringCode = "FFFFFF"
	retryTimes := 5
	// -- WHEN --
	err := Retry(retryTimes, mr.myFunction)
	// -- THEN --
	assert.NotNil(t, err)
}

func resetMonkeyPatching(_ *testing.T) {
	_maxSleepTimeMillis = 10
}
