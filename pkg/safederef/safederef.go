package safederef

func DerefStr(ptr *string) string {
	var result string
	if ptr != nil {
		result = *ptr
	}
	return result
}
