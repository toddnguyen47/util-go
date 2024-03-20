package mathutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GivenMinA_WhenMin_ThenReturnProperMin(t *testing.T) {
	// -- GIVEN --
	a := 40
	b := 42
	// -- WHEN --
	min1 := MinInt(a, b)
	// -- THEN --
	assert.Equal(t, a, min1)
}

func Test_GivenMinB_WhenMin_ThenReturnProperMin(t *testing.T) {
	// -- GIVEN --
	a := 42
	b := 40
	// -- WHEN --
	min1 := MinInt(a, b)
	// -- THEN --
	assert.Equal(t, b, min1)
}

func Test_GivenEqual_WhenMin_ThenReturnProperMin(t *testing.T) {
	// -- GIVEN --
	a := 42
	b := 42
	// -- WHEN --
	min1 := MinInt(a, b)
	// -- THEN --
	assert.Equal(t, b, min1)
}

func Test_GivenMaxB_WhenMax_ThenReturnProperMax(t *testing.T) {
	// -- GIVEN --
	a := 40
	b := 42
	// -- WHEN --
	max1 := MaxInt(a, b)
	// -- THEN --
	assert.Equal(t, b, max1)
}

func Test_GivenMaxA_WhenMax_ThenReturnProperMax(t *testing.T) {
	// -- GIVEN --
	a := 42
	b := 40
	// -- WHEN --
	max1 := MaxInt(a, b)
	// -- THEN --
	assert.Equal(t, a, max1)
}

func Test_GivenEqual_WhenMax_ThenReturnProperMax(t *testing.T) {
	// -- GIVEN --
	a := 42
	b := 42
	// -- WHEN --
	max1 := MaxInt(a, b)
	// -- THEN --
	assert.Equal(t, a, max1)
}

func Test_GivenDenominatorNotZero_WhenDivide_ThenDivideProperly(t *testing.T) {
	// -- GIVEN --
	numerator := 10
	denominator := 4
	// -- WHEN --
	quotient := DivideSafely(float64(numerator), float64(denominator))
	// -- THEN --
	assert.Equal(t, 2.5, quotient)
}

func Test_GivenDenominatorZero_WhenDivide_ThenReturn0(t *testing.T) {
	// -- GIVEN --
	numerator := 0
	denominator := 0
	// -- WHEN --
	quotient := DivideSafely(float64(numerator), float64(denominator))
	// -- THEN --
	assert.Equal(t, 0.0, quotient)
}
