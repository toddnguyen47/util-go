package retryjitter

import (
	"context"
	"crypto/rand"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _contextForTests = context.Background()
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
	// -- ARRANGE --
	resetMonkeyPatching(t)
	mr := new(mockRetry)
	mr.stringCode = "FFP"
	retryTimes := 3
	// -- ACT --
	err := Retry(_contextForTests, retryTimes, mr.myFunction)
	// -- ASSERT --
	assert.Nil(t, err)
}

func Test_GivenRetrySuccessButTimeoutLessThanZero_ThenTimeoutDefaultsTo100ReturnErrorNil(t *testing.T) {
	// -- ARRANGE --
	resetMonkeyPatching(t)
	mr := new(mockRetry)
	mr.stringCode = "FFP"
	retryTimes := 3
	// -- ACT --
	err := RetryWithTimeout(_contextForTests, retryTimes, -500, mr.myFunction)
	// -- ASSERT --
	assert.Nil(t, err)
}

func Test_GivenRetrySuccessButMinSleepTimeGreaterThanMaxSleepTime_ThenTimeoutDefaultsTo100ReturnErrorNil(t *testing.T) {
	// -- ARRANGE --
	resetMonkeyPatching(t)
	mr := new(mockRetry)
	mr.stringCode = "FFP"
	retryTimes := 3
	_minSleepTimeMillis = 50
	// -- ACT --
	err := RetryWithTimeout(_contextForTests, retryTimes, 1, mr.myFunction)
	// -- ASSERT --
	assert.Nil(t, err)
}

func Test_GivenGeneratingRandomIntErrorRetrySuccess_ThenReturnNil(t *testing.T) {
	// -- ARRANGE --
	_reader = errReadCloser(1)
	mr := new(mockRetry)
	mr.stringCode = "FFP"
	retryTimes := 3
	// -- ACT --
	err := Retry(_contextForTests, retryTimes, mr.myFunction)
	// -- ASSERT --
	assert.Nil(t, err)
}

func Test_GivenRetryFailure_ThenReturnErr(t *testing.T) {
	resetMonkeyPatching(t)
	// -- ARRANGE --
	mr := new(mockRetry)
	mr.stringCode = "FFFFFF"
	retryTimes := 5
	// -- ACT --
	err := Retry(_contextForTests, retryTimes, mr.myFunction)
	// -- ASSERT --
	assert.NotNil(t, err)
}

func resetMonkeyPatching(_ *testing.T) {
	_reader = rand.Reader
}

// ------------------------------------------------------------
// #region errReadCloser

// errReadCloser - Ref: https://stackoverflow.com/a/45126402/6323360
type errReadCloser int

func (m errReadCloser) Read(_ []byte) (n int, err error) {
	return 0, errors.New("some error")
}

func (m errReadCloser) Close() error {
	return nil
}

// #endregion
// o----------------------------------------------------------o
