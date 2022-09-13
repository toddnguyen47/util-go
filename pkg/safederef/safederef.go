package safederef

func Deref[T comparable](ptr *T) T {
	var result T
	if ptr != nil {
		result = *ptr
	}
	return result
}

func DerefSlice[T comparable](ptr *[]T) []T {
	result := make([]T, 0)
	if ptr != nil {
		result = *ptr
	}
	return result
}
