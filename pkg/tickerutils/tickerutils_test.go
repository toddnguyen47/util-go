package tickerutils

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// /@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\
// #region SETUP
// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type TickerUtilsTestSuite struct {
	suite.Suite
	ctxBg context.Context
	count atomic.Uint32
}

func (s *TickerUtilsTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
	s.count = atomic.Uint32{}
}

func (s *TickerUtilsTestSuite) TearDownTest() {
	s.resetMonkeyPatching()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestTickerUtilsTestSuite(t *testing.T) {
	suite.Run(t, new(TickerUtilsTestSuite))
}

// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// #endregion SETUP
// \@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@/

// /@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\
// #region TESTS ARE BELOW
// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

func (s *TickerUtilsTestSuite) Test_GivenTickerTimedOut_ThenFunctionIsCalled() {
	// -- ARRANGE --
	sleepTime := 10 * time.Millisecond
	sutTicker2 := NewTicker2(sleepTime, s.testFuncIdle)
	// -- ACT --
	time.Sleep(sleepTime * 5)
	// -- ASSERT --
	sutTicker2.Stop()
	// Stop twice on purpose
	sutTicker2.Stop()
	assert.Equal(s.T(), 5, int(s.count.Load()))
}

// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// #endregion TESTS
// \@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@/

// /@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\
// #region TEST HELPERS
// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

func (s *TickerUtilsTestSuite) resetMonkeyPatching() {
}

func (s *TickerUtilsTestSuite) testFuncIdle(t1 time.Time) {
	s.count.Add(1)
	fmt.Println("In testFuncIdle()", t1.String())
}

// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// #region TEST HELPERS
// \@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@/
