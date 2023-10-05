package retryclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// ############################################################################
// #region SETUP
// ############################################################################

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type RetryClientTestSuite struct {
	suite.Suite
	ctxBg      context.Context
	mockClient *mockClientStruct
}

func (s *RetryClientTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
	s.mockClient = new(mockClientStruct)
	s.mockClient.t = s.T()
}

func (s *RetryClientTestSuite) TearDownTest() {
	s.resetMonkeyPatching()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestRetryClientTestSuite(t *testing.T) {
	suite.Run(t, new(RetryClientTestSuite))
}

// #endregion

// ############################################################################
// #region TESTS ARE BELOW
// ############################################################################

func (s *RetryClientTestSuite) Test_GivenRetrySuccessfulWith1Retry_ThenErrIsNil() {
	// -- ARRANGE --
	req, err := http.NewRequestWithContext(s.ctxBg, http.MethodGet, "url",
		io.NopCloser(bytes.NewReader([]byte("{}"))))
	assert.Nil(s.T(), err)
	// -- ACT --
	resp, err := Retry(RetryConfig{
		RetryTimes: 1,
		SleepTime:  1 * time.Millisecond,
		Client:     s.mockClient,
		Request:    req,
	})
	// -- ASSERT --
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), http.StatusOK, resp.StatusCode)
}

func (s *RetryClientTestSuite) Test_GivenRetrySuccessfulWith3Retries_ThenErrIsNil() {
	// -- ARRANGE --
	req, err := http.NewRequestWithContext(s.ctxBg, http.MethodGet, "url",
		io.NopCloser(bytes.NewReader([]byte("{}"))))
	assert.Nil(s.T(), err)
	// -- ACT --
	resp, err := Retry(RetryConfig{
		RetryTimes: 3,
		SleepTime:  1 * time.Millisecond,
		Client:     s.mockClient,
		Request:    req,
	})
	// -- ASSERT --
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), http.StatusOK, resp.StatusCode)
}

func (s *RetryClientTestSuite) Test_GivenRetryUnsuccessful_ThenErrIsNotNil() {
	// -- ARRANGE --
	req, err := http.NewRequestWithContext(s.ctxBg, http.MethodGet, "url", nil)
	assert.Nil(s.T(), err)
	retryTimes := uint(3)
	s.mockClient.statusCode = http.StatusBadRequest
	// -- ACT --
	resp, err := Retry(RetryConfig{
		RetryTimes: retryTimes,
		SleepTime:  1 * time.Millisecond,
		Client:     s.mockClient,
		Request:    req,
	})
	// -- ASSERT --
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), http.StatusBadRequest, resp.StatusCode)
}

func (s *RetryClientTestSuite) Test_GivenClientErr_ThenReturnClientErr() {
	// -- ARRANGE --
	req, err := http.NewRequestWithContext(s.ctxBg, http.MethodGet, "url",
		io.NopCloser(bytes.NewReader([]byte("{}"))))
	assert.Nil(s.T(), err)
	s.mockClient.numDoErrCount = 5
	// -- ACT --
	resp, err := Retry(RetryConfig{
		RetryTimes: 3,
		SleepTime:  1 * time.Millisecond,
		Client:     s.mockClient,
		Request:    req,
	})
	// -- ASSERT --
	assert.Equal(s.T(), err, errForTests)
	assert.Equal(s.T(), http.StatusInternalServerError, resp.StatusCode)
}

func (s *RetryClientTestSuite) Test_GivenRetryLessThanZero_ThenErrIsNotNil() {
	// -- ARRANGE --
	req, err := http.NewRequestWithContext(s.ctxBg, http.MethodGet, "url", nil)
	assert.Nil(s.T(), err)
	// -- ACT --
	resp, err := Retry(RetryConfig{
		RetryTimes: 0,
		SleepTime:  1 * time.Millisecond,
		Client:     s.mockClient,
		Request:    req,
	})
	// -- ASSERT --
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), http.StatusInternalServerError, resp.StatusCode)
}

func (s *RetryClientTestSuite) Test_GivenIoReadAllError_ThenErrIsNotNil() {
	// -- ARRANGE --
	req, err := http.NewRequestWithContext(s.ctxBg, http.MethodGet, "url", errReadCloser(1))
	assert.Nil(s.T(), err)
	// -- ACT --
	resp, err := Retry(RetryConfig{
		RetryTimes: 3,
		SleepTime:  1 * time.Millisecond,
		Client:     s.mockClient,
		Request:    req,
	})
	// -- ASSERT --
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), http.StatusInternalServerError, resp.StatusCode)
}

// ############################################################################
// #region TEST HELPERS
// ############################################################################

func (s *RetryClientTestSuite) resetMonkeyPatching() {
}

// #endregion

// mockClientStruct - mocking HTTP Client
type mockClientStruct struct {
	numDoCount    uint32
	numDoErrCount uint32
	statusCode    int
	t             *testing.T
}

func (m *mockClientStruct) Do(req *http.Request) (*http.Response, error) {
	atomic.AddUint32(&m.numDoCount, 1)
	if atomic.LoadUint32(&m.numDoErrCount) > 0 {
		atomic.AddUint32(&m.numDoErrCount, ^uint32(0))
		resp := http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
		}
		return &resp, errForTests
	}

	// Read body
	if req.Body != nil {
		b1, err := io.ReadAll(req.Body)
		assert.Nil(m.t, err)
		var output interface{}
		err = json.Unmarshal(b1, &output)
		assert.Nil(m.t, err)
		err = req.Body.Close()
		assert.Nil(m.t, err)
	}

	statusCode := http.StatusOK
	if m.statusCode != 0 {
		statusCode = m.statusCode
	}
	resp := http.Response{
		StatusCode: statusCode,
		Body:       getHttpBody("{}"),
	}

	return &resp, nil
}

func getHttpBody(input string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(input)))
}

var errForTests = errors.New("test error")

// ------------------------------------------------------------
// #region errReadCloser

// errReadCloser - Ref: https://stackoverflow.com/a/45126402/6323360
type errReadCloser int

func (m errReadCloser) Read(_ []byte) (n int, err error) {
	return 0, errors.New("some error")
}

func (m errReadCloser) Close() error {
	return nil
}

// #endregion
// o----------------------------------------------------------o
