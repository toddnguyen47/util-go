package listuniqueitems

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GivenDuplicateItemsList_When_ThenItemsReturnedAreUnique(t *testing.T) {
	list1 := []int{42, 42, 5, 5, 5, 1, 26, 26}

	newList := GetListOfUniqueItems(list1)

	expectedList := []int{42, 5, 1, 26}
	assert.Equal(t, expectedList, newList)
}

func Test_GivenNoDuplicateItemsList_When_ThenItemsReturnedAreUnique(t *testing.T) {
	list1 := []int{42, 5, 1, 26}

	newList := GetListOfUniqueItems(list1)

	expectedList := []int{42, 5, 1, 26}
	assert.Equal(t, expectedList, newList)
}

func Test_GivenNilList_When_ThenItemsReturnedAreEmpty(t *testing.T) {
	var list1 []int

	newList := GetListOfUniqueItems(list1)

	var expectedList []int
	assert.Equal(t, expectedList, newList)
}

func Test_GivenEmptyList_When_ThenItemsReturnedAreEmpty(t *testing.T) {
	list1 := make([]int, 0)

	newList := GetListOfUniqueItems(list1)

	expectedList := make([]int, 0)
	assert.Equal(t, expectedList, newList)
}
