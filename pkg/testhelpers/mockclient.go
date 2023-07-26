package testhelpers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ------------------------------------------------------------
// #region mockClientStruct

type clientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

type MockClient interface {
	clientInterface

	MpfDo() MockPassFail
	SavedRequest() *http.Request
	ExpectedResponse() *http.Response
	SetExpectedResponse(response *http.Response)
}

func NewMockClient(t *testing.T) MockClient {
	m := mockClientStruct{
		mpfDo:        NewMockPassFail(),
		savedRequest: nil,
		expectedResp: nil,
		t:            t,
		mutex:        sync.Mutex{},
	}
	return &m
}

// mockClientStruct - mocking HTTP Client
type mockClientStruct struct {
	mpfDo        MockPassFail
	savedRequest *http.Request
	expectedResp *http.Response
	t            *testing.T
	mutex        sync.Mutex
}

func (m *mockClientStruct) Do(req *http.Request) (*http.Response, error) {
	m.mutex.Lock()
	m.savedRequest = req
	m.mutex.Unlock()

	err := m.mpfDo.WillPassIncrementCount()
	if err != nil {
		return &http.Response{}, ErrFunctionShouldFail
	}

	// Read body
	if req.Body != nil {
		b1, err2 := io.ReadAll(req.Body)
		assert.Nil(m.t, err2)
		var output interface{}
		err2 = json.Unmarshal(b1, &output)
		assert.Nil(m.t, err2)
		err2 = req.Body.Close()
		assert.Nil(m.t, err2)
	}

	var resp *http.Response
	if m.expectedResp == nil {
		resp = &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader("{}")),
		}
	} else {
		resp = m.expectedResp
	}

	return resp, nil
}

func (m *mockClientStruct) MpfDo() MockPassFail { return m.mpfDo }

func (m *mockClientStruct) SavedRequest() *http.Request { return m.savedRequest }

func (m *mockClientStruct) ExpectedResponse() *http.Response { return m.expectedResp }

func (m *mockClientStruct) SetExpectedResponse(response *http.Response) { m.expectedResp = response }
