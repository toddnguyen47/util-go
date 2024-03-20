package mathutils

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func DivideSafely(numerator, denominator float64) float64 {
	if denominator == 0 {
		return 0.00
	}
	return numerator / denominator
}
