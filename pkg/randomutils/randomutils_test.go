package randomutils

import (
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toddnguyen47/util-go/pkg/testhelpers"
)

func Test_GivenRandomNumber_ThenNumberIsBetweenRange(t *testing.T) {
	// -- GIVEN --
	minNumber := int64(50)
	maxNumber := int64(200)
	_reader = rand.Reader
	// Test 10 times
	for i := 0; i < 100; i++ {
		// -- WHEN --
		randomNumber := GetRandomWithMin(maxNumber, minNumber)
		// -- THEN --
		assert.True(t, minNumber <= randomNumber && randomNumber <= maxNumber,
			fmt.Sprintf("failed -> min: %d, max: %d, random: %d", minNumber, maxNumber, randomNumber))
	}
}

func Test_GivenMinMaxIsTheSame_ThenNumberIsBetweenRange(t *testing.T) {
	// -- GIVEN --
	minNumber := int64(50)
	maxNumber := int64(50)
	_reader = rand.Reader
	// Test 10 times
	for i := 0; i < 100; i++ {
		// -- WHEN --
		randomNumber := GetRandomWithMin(maxNumber, minNumber)
		// -- THEN --
		assert.True(t, minNumber <= randomNumber && randomNumber <= maxNumber,
			fmt.Sprintf("failed -> min: %d, max: %d, random: %d", minNumber, maxNumber, randomNumber))
	}
}

func Test_GivenRandIntError_ThenUseMaxRange(t *testing.T) {
	// -- GIVEN --
	minNumber := int64(50)
	maxNumber := int64(200)
	_reader = testhelpers.ErrReadCloser(-1)
	// Test 10 times
	for i := 0; i < 100; i++ {
		// -- WHEN --
		randomNumber := GetRandomWithMin(maxNumber, minNumber)
		// -- THEN --
		assert.True(t, minNumber <= randomNumber && randomNumber <= maxNumber,
			fmt.Sprintf("failed -> min: %d, max: %d, random: %d", minNumber, maxNumber, randomNumber))
	}
}

func Test_GivenMinMaxBothZero_ThenUse1AsMax(t *testing.T) {
	// -- GIVEN --
	minNumber := int64(0)
	maxNumber := int64(0)
	_reader = rand.Reader
	// Test 10 times
	for i := 0; i < 100; i++ {
		// -- WHEN --
		randomNumber := GetRandomWithMin(maxNumber, minNumber)
		// -- THEN --
		assert.True(t, minNumber <= randomNumber && randomNumber <= maxNumber,
			fmt.Sprintf("failed -> min: %d, max: %d, random: %d", minNumber, maxNumber, randomNumber))
	}
}
