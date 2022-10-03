package paginate

// Paginate - paginationSize will have a minimum of 1
func Paginate[T interface{}](listInput []T, paginationSize int) [][]T {
	results := make([][]T, 0)

	if paginationSize < 1 {
		paginationSize = 1
	}

	curResult := make([]T, 0)
	len1 := len(listInput)
	for i, elem := range listInput {
		curResult = append(curResult, elem)

		nextIndex := i + 1
		if nextIndex == len1 || nextIndex%paginationSize == 0 {
			results = append(results, curResult)
			// Clear
			curResult = make([]T, 0)
		}
	}

	return results
}
