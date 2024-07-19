package mapequalfold

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

// /--------------------------------------------------------------------------\
// #region SETUP
// ----------------------------------------------------------------------------

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type MapEqualFoldTestSuite struct {
	suite.Suite
	ctxBg context.Context
}

func (s *MapEqualFoldTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
}

func (s *MapEqualFoldTestSuite) TearDownTest() {
	s.resetMonkeyPatching()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestMapEqualFoldTestSuite(t *testing.T) {
	suite.Run(t, new(MapEqualFoldTestSuite))
}

// ----------------------------------------------------------------------------
// #endregion SETUP
// \--------------------------------------------------------------------------/

// /--------------------------------------------------------------------------\
// #region TESTS ARE BELOW
// ----------------------------------------------------------------------------

func (s *MapEqualFoldTestSuite) Test_MapEmptyStruct() {
	// -- GIVEN --
	map1 := NewMapEqualFold[*EmptyStruct]()
	map1.Set("hElLo", &EmptyStruct{})
	// -- WHEN --
	es, found := map1.Get("HELLO")
	// -- THEN --
	s.NotNil(es)
	s.True(found)
	s.Equal("", es.String())
	// -- ASSERT AGAIN --
	es2, found2 := map1.Get("helo")
	s.Nil(es2)
	s.False(found2)
}

// ----------------------------------------------------------------------------
// #endregion TESTS``
// \--------------------------------------------------------------------------/

// /--------------------------------------------------------------------------\
// #region TEST HELPERS
// ----------------------------------------------------------------------------

func (s *MapEqualFoldTestSuite) resetMonkeyPatching() {
}

// ----------------------------------------------------------------------------
// #region TEST HELPERS
// \--------------------------------------------------------------------------/
