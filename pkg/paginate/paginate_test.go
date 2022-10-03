package paginate

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// ############################################################################
// #region SETUP
// ############################################################################

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

// ############################################################################
// #region TESTS ARE BELOW
// ############################################################################

func (s *PaginateTestSuite) Test_Given8GroupsPaginationSize4_When_ThenReturn2Groups() {
	// -- ARRANGE --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60}
	// -- ACT --
	newList := Paginate(list1, 4)
	// -- ASSERT --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
	}
	assert.Equal(s.T(), expectedResults, newList)
}

func (s *PaginateTestSuite) Test_Given9GroupsPaginationSize4_When_ThenReturn3Groups() {
	// -- ARRANGE --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99}
	// -- ACT --
	newList := Paginate(list1, 4)
	// -- ASSERT --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
		{99},
	}
	assert.Equal(s.T(), expectedResults, newList)
}

func (s *PaginateTestSuite) Test_Given10GroupsPaginationSize4_When_ThenReturn3Groups() {
	// -- ARRANGE --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74}
	// -- ACT --
	newList := Paginate(list1, 4)
	// -- ASSERT --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
		{99, 74},
	}
	assert.Equal(s.T(), expectedResults, newList)
}

func (s *PaginateTestSuite) Test_Given11GroupsPaginationSize4_When_ThenReturn3Groups() {
	// -- ARRANGE --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74, -9}
	// -- ACT --
	newList := Paginate(list1, 4)
	// -- ASSERT --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
		{99, 74, -9},
	}
	assert.Equal(s.T(), expectedResults, newList)
}

func (s *PaginateTestSuite) Test_Given12GroupsPaginationSize4_When_ThenReturn3Groups() {
	// -- ARRANGE --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74, -9, 105}
	// -- ACT --
	newList := Paginate(list1, 4)
	// -- ASSERT --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
		{99, 74, -9, 105},
	}
	assert.Equal(s.T(), expectedResults, newList)
}

func (s *PaginateTestSuite) Test_Given13GroupsPaginationSize4_When_ThenReturn4Groups() {
	// -- ARRANGE --
	list1 := []int{5, 100, 42, 11, 4, -1, 16, 60, 99, 74, -9, 105, 44}
	// -- ACT --
	newList := Paginate(list1, 4)
	// -- ASSERT --
	expectedResults := [][]int{
		{5, 100, 42, 11},
		{4, -1, 16, 60},
		{99, 74, -9, 105},
		{44},
	}
	assert.Equal(s.T(), expectedResults, newList)
}

func (s *PaginateTestSuite) Test_Given3GroupsPaginationSizeNeg1_When_ThenReturn4Groups() {
	// -- ARRANGE --
	list1 := []int{5, 100, 42}
	// -- ACT --
	newList := Paginate(list1, -1)
	// -- ASSERT --
	expectedResults := [][]int{
		{5}, {100}, {42},
	}
	assert.Equal(s.T(), expectedResults, newList)
}

// ############################################################################
// #region TEST HELPERS
// ############################################################################

func (s *PaginateTestSuite) resetMonkeyPatching() {
}

// #endregion
