package paginate

// SimplePaginate - paginationSize will have a minimum of 1
// Paginate items until the batch hits `paginationSize`
func SimplePaginate[T interface{}](listInput []T, paginationSize int) [][]T {

	results := make([][]T, 0)
	len1 := len(listInput)
	if len1 == 0 {
		return results
	}

	if paginationSize < 1 {
		paginationSize = 1
	}

	curResult := make([]T, 0)
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

// EvenPaginate - Paginate so that batches have even number of items, up to a max of
// `paginateSize` items. For example, a group of 14 items, pagination size of 4 will
// have the number of items [4, 4, 3, 3] instead of [4, 4, 4, 2]
func EvenPaginate[T interface{}](listInput []T, paginationSize int) [][]T {

	results := make([][]T, 0)
	lenListInput := len(listInput)
	if lenListInput == 0 {
		return results
	}

	if paginationSize < 1 {
		paginationSize = 1
	}
	numberOfBucketsNeeded := lenListInput / paginationSize
	if lenListInput%paginationSize != 0 {
		numberOfBucketsNeeded += 1
	}

	minItems := lenListInput / numberOfBucketsNeeded
	maxItems := minItems + 1
	bucketsWithMaxItems := lenListInput % numberOfBucketsNeeded
	getBucketSize := func(index int) int {
		if index < bucketsWithMaxItems {
			return maxItems
		}
		return minItems
	}
	curResult := make([]T, 0)
	currentBucket := 0
	currentBucketSize := getBucketSize(currentBucket)

	for i, elem := range listInput {
		curResult = append(curResult, elem)
		if i == lenListInput-1 || len(curResult) >= currentBucketSize {
			results = append(results, curResult)
			// Clear
			curResult = make([]T, 0)
			currentBucket += 1
			currentBucketSize = getBucketSize(currentBucket)
		}
	}

	return results
}
