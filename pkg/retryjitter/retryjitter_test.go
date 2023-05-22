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
	// -- ARRANGE --
	mr := new(mockRetry)
	mr.stringCode = "FFP"
	retryTimes := 3
	// -- ACT --
	err := Retry(retryTimes, mr.myFunction)
	// -- ASSERT --
	assert.Nil(t, err)
}

func Test_GivenRetryFailure_ThenReturnErr(t *testing.T) {
	// -- ARRANGE --
	mr := new(mockRetry)
	mr.stringCode = "FFFFFF"
	retryTimes := 5
	// -- ACT --
	err := Retry(retryTimes, mr.myFunction)
	// -- ASSERT --
	assert.NotNil(t, err)
}
