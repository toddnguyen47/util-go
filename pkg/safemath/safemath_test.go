package safemath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GivenDenomZero_ThenReturnZero(t *testing.T) {
	numerator := 42.00
	denominator := 0.0000000000

	quotient := Divide(numerator, denominator)

	assert.Equal(t, 0.00, quotient)
}

func Test_GivenDenomNonZero_ThenReturnDivision(t *testing.T) {
	numerator := 42.00
	denominator := 0.0000000001

	quotient := Divide(numerator, denominator)

	assert.Equal(t, numerator/denominator, quotient)
}
