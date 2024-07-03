package testhelpers

import (
	"context"
	"io"
	"net/http"
	"strings"
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
type MockClientTestSuite struct {
	suite.Suite
	ctxBg context.Context
}

func (s *MockClientTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
}

func (s *MockClientTestSuite) TearDownTest() {
	s.resetMonkeyPatching()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestMockClientTestSuite(t *testing.T) {
	suite.Run(t, new(MockClientTestSuite))
}

// #endregion

// ------------------------------------------------------------
// #region TESTS ARE BELOW
// ------------------------------------------------------------

func (s *MockClientTestSuite) Test_GivenNewMockClient_ThenReturnMockClient() {
	// -- GIVEN --
	mockClient := NewMockClient(s.T())
	req, err := http.NewRequestWithContext(s.ctxBg, http.MethodGet, "url", strings.NewReader("{}"))
	assert.Nil(s.T(), err)
	// -- WHEN --
	resp, err := mockClient.Do(req)
	// -- THEN --
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *MockClientTestSuite) Test_GivenSetCodeFail_ThenReturnError() {
	// -- GIVEN --
	mockClient := NewMockClient(s.T())
	req, err := http.NewRequestWithContext(s.ctxBg, http.MethodGet, "url", strings.NewReader("{}"))
	assert.Nil(s.T(), err)
	mockClient.MpfDo().SetCode("FFP")
	// -- WHEN --
	resp, err := mockClient.Do(req)
	// -- THEN --
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), &http.Response{}, resp)
	assert.NotNil(s.T(), mockClient.SavedRequest())
}

func (s *MockClientTestSuite) Test_GivenExpectedResp_ThenReturnResp() {
	// -- GIVEN --
	mockClient := NewMockClient(s.T())
	req, err := http.NewRequestWithContext(s.ctxBg, http.MethodGet, "url", strings.NewReader("{}"))
	assert.Nil(s.T(), err)
	expectedResp := http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader("{}")),
	}
	mockClient.SetExpectedResponse(&expectedResp)
	// -- WHEN --
	resp, err := mockClient.Do(req)
	// -- THEN --
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), &expectedResp, resp)
	assert.Equal(s.T(), &expectedResp, mockClient.ExpectedResponse())
}

// ------------------------------------------------------------
// #region TEST HELPERS
// ------------------------------------------------------------

func (s *MockClientTestSuite) resetMonkeyPatching() {
}

// #endregion
