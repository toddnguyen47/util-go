package paginate

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// ------------------------------------------------------------
// #region SETUP
// ------------------------------------------------------------

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type PaginateTestSuite struct {
	suite.Suite
	ctxBg context.Context
}

func (s *PaginateTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
}

func (s *PaginateTestSuite) TearDownTest() {
	s.resetMonkeyPatching()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPaginateTestSuite(t *testing.T) {
	suite.Run(t, new(PaginateTestSuite))
}

// #endregion

// ------------------------------------------------------------
// #region TESTS ARE BELOW
// ------------------------------------------------------------

func (s *PaginateTestSuite) Test_Given8ElemsPaginationSize4_WhenSimplePaginate_ThenReturn2Groups() {
	// -- GIVEN --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60}
	// -- WHEN --
	newList := SimplePaginate(list1, 4)
	// -- THEN --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
	}
	assert.Equal(s.T(), expectedResults, newList)
}

func (s *PaginateTestSuite) Test_Given9ElemsPaginationSize4_When_ThenReturn3Groups() {
	// -- GIVEN --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99}
	// -- WHEN --
	newList := SimplePaginate(list1, 4)
	// -- THEN --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
		{99},
	}
	assert.Equal(s.T(), expectedResults, newList)
}

func (s *PaginateTestSuite) Test_Given10ElemsPaginationSize4_WhenSimplePaginate_ThenReturn3Groups() {
	// -- GIVEN --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74}
	// -- WHEN --
	newList := SimplePaginate(list1, 4)
	// -- THEN --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
		{99, 74},
	}
	assert.Equal(s.T(), expectedResults, newList)
}

func (s *PaginateTestSuite) Test_Given11ElemsPaginationSize4_WhenSimplePaginate_ThenReturn3Groups() {
	// -- GIVEN --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74, -9}
	// -- WHEN --
	newList := SimplePaginate(list1, 4)
	// -- THEN --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
		{99, 74, -9},
	}
	assert.Equal(s.T(), expectedResults, newList)
}

func (s *PaginateTestSuite) Test_Given12ElemsPaginationSize4_WhenSimplePaginate_ThenReturn3Groups() {
	// -- GIVEN --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74, -9, 105}
	// -- WHEN --
	newList := SimplePaginate(list1, 4)
	// -- THEN --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
		{99, 74, -9, 105},
	}
	assert.Equal(s.T(), expectedResults, newList)
}

func (s *PaginateTestSuite) Test_Given13ElemsPaginationSize4_WhenSimplePaginate_ThenReturn4Groups() {
	// -- GIVEN --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74, -9, 105, 44}
	// -- WHEN --
	newList := SimplePaginate(list1, 4)
	// -- THEN --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
		{99, 74, -9, 105},
		{44},
	}
	assert.Equal(s.T(), expectedResults, newList)
}

func (s *PaginateTestSuite) Test_Given3ElemsPaginationSizeNeg1_WhenSimplePaginate_ThenReturn4Groups() {
	// -- GIVEN --
	list1 := []int{5, 100, 42}
	// -- WHEN --
	newList := SimplePaginate(list1, -1)
	// -- THEN --
	expectedResults := [][]int{
		{5}, {100}, {42},
	}
	assert.Equal(s.T(), expectedResults, newList)
}

func (s *PaginateTestSuite) Test_Given9ElemsPaginationSizeNeg1_WhenEvenPaginate_ThenReturn1EvenGroups() {
	// -- GIVEN --
	list1 := []int{5, 100, 42, 11}
	// -- WHEN --
	newList := EvenPaginate(list1, -1)
	// -- THEN --
	expectedResults := [][]int{
		{5}, {100}, {42}, {11},
	}

	assert.Equal(s.T(), expectedResults, newList)
}

func (s *PaginateTestSuite) Test_Given9ElemsPaginationSize4_WhenEvenPaginate_ThenReturn3EvenGroups() {
	// -- GIVEN --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99}
	// -- WHEN --
	newList := EvenPaginate(list1, 4)
	// -- THEN --
	expectedResults := [][]int{
		{5, 100, 42},
		{11, 4, -1},
		{16, 60, 99},
	}
	assert.Equal(s.T(), expectedResults, newList)
}

func (s *PaginateTestSuite) Test_Given10ElemsPaginationSize4_WhenEvenPaginate_ThenReturn3EvenGroups() {
	// -- GIVEN --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74}
	// -- WHEN --
	newList := EvenPaginate(list1, 4)
	// -- THEN --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16},
		{60, 99, 74},
	}
	assert.Equal(s.T(), expectedResults, newList)
}

func (s *PaginateTestSuite) Test_Given11ElemsPaginationSize4_WhenEvenPaginate_ThenReturn3EvenGroups() {
	// -- GIVEN --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74, -9}
	// -- WHEN --
	newList := EvenPaginate(list1, 4)
	// -- THEN --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
		{99, 74, -9},
	}
	assert.Equal(s.T(), expectedResults, newList)
}

func (s *PaginateTestSuite) Test_Given12ElemsPaginationSize4_WhenEvenPaginate_ThenReturn3Groups() {
	// -- GIVEN --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74, -9, 105}
	// -- WHEN --
	newList := SimplePaginate(list1, 4)
	// -- THEN --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
		{99, 74, -9, 105},
	}
	assert.Equal(s.T(), expectedResults, newList)
}

func (s *PaginateTestSuite) Test_Given13ElemsPaginationSize4_WhenEvenPaginate_ThenReturn4Groups() {
	// -- GIVEN --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74, -9, 105, 44}
	// -- WHEN --
	newList := EvenPaginate(list1, 4)
	// -- THEN --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16},
		{60, 99, 74},
		{-9, 105, 44},
	}
	assert.Equal(s.T(), expectedResults, newList)
}

func (s *PaginateTestSuite) Test_Given14ElemsPaginationSize4_WhenEvenPaginate_ThenReturn4Groups() {
	// -- GIVEN --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74, -9, 105, 44, 22}
	// -- WHEN --
	newList := EvenPaginate(list1, 4)
	// -- THEN --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
		{99, 74, -9},
		{105, 44, 22},
	}
	assert.Equal(s.T(), expectedResults, newList)
}

func (s *PaginateTestSuite) Test_Given0Items_WhenEvenPaginate_ThenReturnOneList() {
	// -- GIVEN --
	var list1 []int
	// -- WHEN --
	newList := EvenPaginate(list1, 4)
	// -- THEN --
	expectedResults := make([][]int, 0)
	assert.Equal(s.T(), expectedResults, newList)
}

func (s *PaginateTestSuite) Test_Given0Items_WhenSimplePaginate_ThenReturnOneList() {
	// -- GIVEN --
	var list1 []int
	// -- WHEN --
	newList := SimplePaginate(list1, 4)
	// -- THEN --
	expectedResults := make([][]int, 0)
	assert.Equal(s.T(), expectedResults, newList)
}

// ------------------------------------------------------------
// #region TEST HELPERS
// ------------------------------------------------------------

func (s *PaginateTestSuite) resetMonkeyPatching() {
}

// #endregion
