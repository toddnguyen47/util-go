package listuniqueitems

import mapset "github.com/deckarep/golang-set/v2"

func GetListOfUniqueItems[T comparable](inputList []T) []T {
	if inputList == nil || len(inputList) <= 0 {
		return inputList
	}

	set1 := mapset.NewSet[T]()
	uniqueList := make([]T, 0)
	for _, message := range inputList {
		if set1.Contains(message) {
			continue
		}
		set1.Add(message)
		uniqueList = append(uniqueList, message)
	}
	return uniqueList
}
