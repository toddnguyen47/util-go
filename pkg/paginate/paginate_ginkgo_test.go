package paginate

import (
	"testing"

	. "github.com/onsi/gomega"
)

// ############################################################################
// #region TESTS ARE BELOW
// ############################################################################

func Test_Given8ElemsPaginationSize4_WhenSimplePaginate_ThenReturn2Groups(t *testing.T) {
	// -- ARRANGE --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60}
	// -- ACT --
	newList := SimplePaginate(list1, 4)
	// -- ASSERT --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
	}
	g := NewWithT(t)
	g.Expect(newList).To(Equal(expectedResults))
}

func Test_Given9ElemsPaginationSize4_When_ThenReturn3Groups(t *testing.T) {
	// -- ARRANGE --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99}
	// -- ACT --
	newList := SimplePaginate(list1, 4)
	// -- ASSERT --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
		{99},
	}
	g := NewWithT(t)
	g.Expect(newList).To(Equal(expectedResults))
}

func Test_Given10ElemsPaginationSize4_WhenSimplePaginate_ThenReturn3Groups(t *testing.T) {
	// -- ARRANGE --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74}
	// -- ACT --
	newList := SimplePaginate(list1, 4)
	// -- ASSERT --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
		{99, 74},
	}
	g := NewWithT(t)
	g.Expect(newList).To(Equal(expectedResults))
}

func Test_Given11ElemsPaginationSize4_WhenSimplePaginate_ThenReturn3Groups(t *testing.T) {
	// -- ARRANGE --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74, -9}
	// -- ACT --
	newList := SimplePaginate(list1, 4)
	// -- ASSERT --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
		{99, 74, -9},
	}
	g := NewWithT(t)
	g.Expect(newList).To(Equal(expectedResults))
}

func Test_Given12ElemsPaginationSize4_WhenSimplePaginate_ThenReturn3Groups(t *testing.T) {
	// -- ARRANGE --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74, -9, 105}
	// -- ACT --
	newList := SimplePaginate(list1, 4)
	// -- ASSERT --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
		{99, 74, -9, 105},
	}
	g := NewWithT(t)
	g.Expect(newList).To(Equal(expectedResults))
}

func Test_Given13ElemsPaginationSize4_WhenSimplePaginate_ThenReturn4Groups(t *testing.T) {
	// -- ARRANGE --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74, -9, 105, 44}
	// -- ACT --
	newList := SimplePaginate(list1, 4)
	// -- ASSERT --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
		{99, 74, -9, 105},
		{44},
	}
	g := NewWithT(t)
	g.Expect(newList).To(Equal(expectedResults))
}

func Test_Given3ElemsPaginationSizeNeg1_WhenSimplePaginate_ThenReturn4Groups(t *testing.T) {
	// -- ARRANGE --
	list1 := []int{5, 100, 42}
	// -- ACT --
	newList := SimplePaginate(list1, -1)
	// -- ASSERT --
	expectedResults := [][]int{
		{5}, {100}, {42},
	}
	g := NewWithT(t)
	g.Expect(newList).To(Equal(expectedResults))
}

func Test_Given9ElemsPaginationSizeNeg1_WhenEvenPaginate_ThenReturn1EvenGroups(t *testing.T) {
	// -- ARRANGE --
	list1 := []int{5, 100, 42, 11}
	// -- ACT --
	newList := EvenPaginate(list1, -1)
	// -- ASSERT --
	expectedResults := [][]int{
		{5}, {100}, {42}, {11},
	}

	g := NewWithT(t)
	g.Expect(newList).To(Equal(expectedResults))
}

func Test_Given9ElemsPaginationSize4_WhenEvenPaginate_ThenReturn3EvenGroups(t *testing.T) {
	// -- ARRANGE --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99}
	// -- ACT --
	newList := EvenPaginate(list1, 4)
	// -- ASSERT --
	expectedResults := [][]int{
		{5, 100, 42},
		{11, 4, -1},
		{16, 60, 99},
	}
	g := NewWithT(t)
	g.Expect(newList).To(Equal(expectedResults))
}

func Test_Given10ElemsPaginationSize4_WhenEvenPaginate_ThenReturn3EvenGroups(t *testing.T) {
	// -- ARRANGE --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74}
	// -- ACT --
	newList := EvenPaginate(list1, 4)
	// -- ASSERT --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16},
		{60, 99, 74},
	}
	g := NewWithT(t)
	g.Expect(newList).To(Equal(expectedResults))
}

func Test_Given11ElemsPaginationSize4_WhenEvenPaginate_ThenReturn3EvenGroups(t *testing.T) {
	// -- ARRANGE --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74, -9}
	// -- ACT --
	newList := EvenPaginate(list1, 4)
	// -- ASSERT --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
		{99, 74, -9},
	}
	g := NewWithT(t)
	g.Expect(newList).To(Equal(expectedResults))
}

func Test_Given12ElemsPaginationSize4_WhenEvenPaginate_ThenReturn3Groups(t *testing.T) {
	// -- ARRANGE --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74, -9, 105}
	// -- ACT --
	newList := SimplePaginate(list1, 4)
	// -- ASSERT --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
		{99, 74, -9, 105},
	}
	g := NewWithT(t)
	g.Expect(newList).To(Equal(expectedResults))
}

func Test_Given13ElemsPaginationSize4_WhenEvenPaginate_ThenReturn4Groups(t *testing.T) {
	// -- ARRANGE --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74, -9, 105, 44}
	// -- ACT --
	newList := EvenPaginate(list1, 4)
	// -- ASSERT --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16},
		{60, 99, 74},
		{-9, 105, 44},
	}
	g := NewWithT(t)
	g.Expect(newList).To(Equal(expectedResults))
}

func Test_Given14ElemsPaginationSize4_WhenEvenPaginate_ThenReturn4Groups(t *testing.T) {
	// -- ARRANGE --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74, -9, 105, 44, 22}
	// -- ACT --
	newList := EvenPaginate(list1, 4)
	// -- ASSERT --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
		{99, 74, -9},
		{105, 44, 22},
	}
	g := NewWithT(t)
	g.Expect(newList).To(Equal(expectedResults))
}

func Test_Given0Items_WhenEvenPaginate_ThenReturnOneList(t *testing.T) {
	// -- ARRANGE --
	var list1 []int
	// -- ACT --
	newList := EvenPaginate(list1, 4)
	// -- ASSERT --
	expectedResults := make([][]int, 0)
	g := NewWithT(t)
	g.Expect(newList).To(Equal(expectedResults))
}

func Test_Given0Items_WhenSimplePaginate_ThenReturnOneList(t *testing.T) {
	// -- ARRANGE --
	var list1 []int
	// -- ACT --
	newList := SimplePaginate(list1, 4)
	// -- ASSERT --
	expectedResults := make([][]int, 0)
	g := NewWithT(t)
	g.Expect(newList).To(Equal(expectedResults))
}

// #endregion
