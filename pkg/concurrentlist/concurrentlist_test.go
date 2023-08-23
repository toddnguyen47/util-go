package concurrentlist

import (
	"context"
	"math/rand"
	"sync"
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
type ConcurrentListTestSuite struct {
	suite.Suite
	ctxBg      context.Context
	maxInserts int
}

func (s *ConcurrentListTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
	s.maxInserts = 500
}

func (s *ConcurrentListTestSuite) TearDownTest() {
	s.resetMonkeyPatching()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestConcurrentListTestSuite(t *testing.T) {
	suite.Run(t, new(ConcurrentListTestSuite))
}

// #endregion

// ############################################################################
// #region TESTS ARE BELOW
// ############################################################################

func (s *ConcurrentListTestSuite) Test_GivenConcurrentList_WhenAppendingAndGetting_ThenReturnProperList() {
	// -- ARRANGE --
	l1 := NewConcurrentList[int]()
	var wg sync.WaitGroup
	wg.Add(s.maxInserts)
	// -- ACT --
	for i := 0; i < s.maxInserts; i++ {
		go func() {
			defer wg.Done()
			l1.Append(rand.Intn(1000))
		}()
	}
	wg.Wait()
	// -- ASSERT --
	assert.Equal(s.T(), s.maxInserts, l1.Size())
	assert.False(s.T(), l1.IsEmpty())
	newList := l1.GetList()
	assert.Equal(s.T(), len(newList), l1.Size())
}

func (s *ConcurrentListTestSuite) Test_GivenConcurrentList_WhenGettingBeforeAppending_ThenReturnProperList() {
	// -- ARRANGE --
	l1 := NewConcurrentList[int]()
	l1.Append(42)
	var wg sync.WaitGroup
	wg.Add(s.maxInserts)
	// -- ACT --
	for i := 0; i < s.maxInserts; i++ {
		go func() {
			defer wg.Done()
			val, ok := l1.Get(0)
			assert.Equal(s.T(), 42, val)
			assert.True(s.T(), ok)
			val, ok = l1.Get(-1)
			assert.Equal(s.T(), 0, val)
			assert.False(s.T(), ok)
			l1.Append(rand.Intn(1000))
		}()
	}
	wg.Wait()
	// -- ASSERT --
	assert.Equal(s.T(), s.maxInserts+1, l1.Size())
	assert.False(s.T(), l1.IsEmpty())
	newList := l1.GetList()
	assert.Equal(s.T(), len(newList), l1.Size())
}

func (s *ConcurrentListTestSuite) Test_GivenConcurrentList_WhenUpdating_ThenReturnProperList() {
	// -- ARRANGE --
	l1 := NewConcurrentList[int]()
	l1.Append(42)
	var wg sync.WaitGroup
	wg.Add(s.maxInserts)
	// -- ACT --
	for i := 0; i < s.maxInserts; i++ {
		newRand := rand.Intn(1000)
		go func(i1 int) {
			defer wg.Done()
			ok := l1.Update(0, i1)
			assert.True(s.T(), ok)
			ok = l1.Update(-1, i1)
			assert.False(s.T(), ok)
			l1.Append(rand.Intn(1000))
		}(newRand)
	}
	wg.Wait()
	// -- ASSERT --
	assert.Equal(s.T(), s.maxInserts+1, l1.Size())
	assert.False(s.T(), l1.IsEmpty())
	newList := l1.GetList()
	assert.Equal(s.T(), len(newList), l1.Size())
	val, ok := l1.Get(0)
	assert.NotEqual(s.T(), 42, val)
	assert.True(s.T(), ok)
}

func (s *ConcurrentListTestSuite) Test_GivenConcurrentList_WhenDeleting_ThenReturnProperList() {
	// -- ARRANGE --
	l1 := NewConcurrentList[int]()
	var wg sync.WaitGroup
	wg.Add(s.maxInserts)
	// -- ACT --
	for i := 0; i < s.maxInserts; i++ {
		go func() {
			defer wg.Done()
			l1.Append(rand.Intn(1000))
		}()
	}
	wg.Wait()
	wg.Add(s.maxInserts / 2)
	for i := 0; i < s.maxInserts/2; i++ {
		go func() {
			defer wg.Done()
			l1.Delete(0)
			val, ok := l1.Delete(-1)
			assert.Equal(s.T(), 0, val)
			assert.False(s.T(), ok)
		}()
	}
	wg.Wait()
	// -- ASSERT --
	assert.Equal(s.T(), s.maxInserts/2, l1.Size())
	assert.False(s.T(), l1.IsEmpty())
	newList := l1.GetList()
	assert.Equal(s.T(), len(newList), l1.Size())
	val, ok := l1.Get(0)
	assert.NotEqual(s.T(), 42, val)
	assert.True(s.T(), ok)
}

// ############################################################################
// #region TEST HELPERS
// ############################################################################

func (s *ConcurrentListTestSuite) resetMonkeyPatching() {
}

// #endregion
