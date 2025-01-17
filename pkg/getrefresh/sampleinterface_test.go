package getrefresh

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

// /--------------------------------------------------------------------------\
// #region SETUP
// ----------------------------------------------------------------------------

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type SampleInterfaceTestSuite struct {
	suite.Suite
	ctxBg context.Context
}

func (s *SampleInterfaceTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
}

func (s *SampleInterfaceTestSuite) TearDownTest() {
	s.resetMonkeyPatching()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSampleInterfaceTestSuite(t *testing.T) {
	suite.Run(t, new(SampleInterfaceTestSuite))
}

// ----------------------------------------------------------------------------
// #endregion SETUP
// \--------------------------------------------------------------------------/

// /--------------------------------------------------------------------------\
// #region TESTS ARE BELOW
// ----------------------------------------------------------------------------

func (s *SampleInterfaceTestSuite) Test_GivenGetItemOk_ThenReturnItem() {
	// -- GIVEN --
	getter := NewSampleInterfaceGetter(1 * time.Second)
	// -- WHEN --
	item1 := getter.Get()
	item2 := getter.Get()
	time.Sleep(2 * time.Second)
	item3 := getter.Get()
	item4 := getter.Get()
	// -- THEN --
	s.NotNil(item1)
	s.NotNil(item2)
	s.NotNil(item3)
	s.NotNil(item4)
	s.Equal(item1.SampleFunc(), item2.SampleFunc())
	s.Equal(item1.SampleFunc(), item3.SampleFunc())
	s.Equal(item1.SampleFunc(), item4.SampleFunc())
}

// ----------------------------------------------------------------------------
// #endregion TESTS
// \--------------------------------------------------------------------------/

// /--------------------------------------------------------------------------\
// #region TEST HELPERS
// ----------------------------------------------------------------------------

func (s *SampleInterfaceTestSuite) resetMonkeyPatching() {
}

// ----------------------------------------------------------------------------
// #endregion TEST HELPERS
// \--------------------------------------------------------------------------/
