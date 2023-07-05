package safemath

const _zero64Bit = 0 << 64

func Divide(numerator, denominator float64) float64 {
	// Need to check greater and less than, as equivalent to zero is difficult
	// to compare using bit floating representation.
	if denominator < _zero64Bit || denominator > _zero64Bit {
		return numerator / denominator
	}
	return 0.00
}
