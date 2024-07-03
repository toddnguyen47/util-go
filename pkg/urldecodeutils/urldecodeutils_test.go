package urldecodeutils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

// NOTE: You only need ONE of the `RunSpecs` function in your whole package / suite!
func TestUrlUtilsTestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UrlDecodeUtils Test Suite")
}

var _ = Describe("UrlDecodeUtils Test Suite", func() {

	var ctxBg context.Context
	var mockClient *mockClientStruct

	var mockClientOk = func(statusCode int, respBody string) {
		mockClient.On("Do", mock.Anything).
			Return(
				&http.Response{
					StatusCode: statusCode,
					Body:       io.NopCloser(bytes.NewReader([]byte(respBody))),
				},
				nil,
			)
	}

	BeforeEach(func() {
		resetMonkeyPatching()
		ctxBg = context.Background()
		mockClient = new(mockClientStruct)
	})

	Describe("NewRequestWithContextRouteOffer tests", func() {
		When("Valid call", func() {
			It("returns no error", func() {
				// -- GIVEN --
				mockClientOk(http.StatusOK, "{}")
				req, err := http.NewRequestWithContext(ctxBg, http.MethodPost, "url",
					bytes.NewReader([]byte("{}")))
				Expect(err).To(BeNil())
				var output json.RawMessage
				// -- WHEN --
				statusCode, err := SendRequestDecodeResponse(ctxBg, mockClient, req, &output)
				// -- THEN --
				Expect(err).To(BeNil())
				Expect(output).To(Not(BeNil()))
				Expect(statusCode).To(Equal(http.StatusOK))
			})
		})

		When("Valid call returns 205", func() {
			It("returns no error", func() {
				// -- GIVEN --
				mockClientOk(http.StatusPartialContent, "{}")
				req, err := http.NewRequestWithContext(ctxBg, http.MethodPost, "url",
					bytes.NewReader([]byte("{}")))
				Expect(err).To(BeNil())
				var output json.RawMessage
				// -- WHEN --
				statusCode, err := SendRequestDecodeResponse(ctxBg, mockClient, req, &output)
				// -- THEN --
				Expect(err).To(BeNil())
				Expect(output).To(Not(BeNil()))
				Expect(statusCode).To(Equal(http.StatusPartialContent))
			})
		})

		When("client returns a non 200", func() {
			It("returns error", func() {
				// -- GIVEN --
				mockClientOk(http.StatusBadRequest, "{}")
				req, err := http.NewRequestWithContext(ctxBg, http.MethodPost, "url",
					bytes.NewReader([]byte("{}")))
				Expect(err).To(BeNil())
				var output json.RawMessage
				// -- WHEN --
				statusCode, err := SendRequestDecodeResponse(ctxBg, mockClient, req, &output)
				// -- THEN --
				Expect(err).To(Not(BeNil()))
				Expect(output).To(BeNil())
				Expect(statusCode).To(Equal(http.StatusBadRequest))
			})
		})

		When("Client Do error", func() {
			It("returns error", func() {
				// -- GIVEN --
				mockClient.numDoCount = 2
				req, err := http.NewRequestWithContext(ctxBg, http.MethodPost, "url",
					bytes.NewReader([]byte("{}")))
				Expect(err).To(BeNil())
				var output json.RawMessage
				// -- WHEN --
				statusCode, err := SendRequestDecodeResponse(ctxBg, mockClient, req, &output)
				// -- THEN --
				Expect(err).To(Not(BeNil()))
				Expect(output).To(BeNil())
				Expect(statusCode).To(Equal(http.StatusBadRequest))
			})
		})

		When("IO Read All error on response", func() {
			It("returns error", func() {
				// -- GIVEN --
				mockClient.On("Do", mock.Anything).
					Return(
						&http.Response{
							StatusCode: http.StatusOK,
							Body:       errReadCloser(1),
						},
						nil,
					)
				req, err := http.NewRequestWithContext(ctxBg, http.MethodPost, "url",
					bytes.NewReader([]byte("{}")))
				Expect(err).To(BeNil())
				var output json.RawMessage
				// -- WHEN --
				statusCode, err := SendRequestDecodeResponse(ctxBg, mockClient, req, &output)
				// -- THEN --
				Expect(err).To(Not(BeNil()))
				Expect(output).To(BeNil())
				Expect(statusCode).To(Equal(http.StatusInternalServerError))
			})
		})

		When("JSON Unmarshal error", func() {
			It("returns error", func() {
				// -- GIVEN --
				mockClientOk(http.StatusOK, "{}")
				req, err := http.NewRequestWithContext(ctxBg, http.MethodPost, "url",
					bytes.NewReader([]byte("{}")))
				Expect(err).To(BeNil())
				var output chan string
				// -- WHEN --
				statusCode, err := SendRequestDecodeResponse(ctxBg, mockClient, req, &output)
				// -- THEN --
				Expect(err).To(Not(BeNil()))
				Expect(output).To(BeNil())
				Expect(statusCode).To(Equal(http.StatusInternalServerError))
			})
		})
	})

	Describe("DecodeRequestBody tests", func() {
		When("valid", func() {
			It("returns no error", func() {
				// -- GIVEN --
				req, err := http.NewRequestWithContext(ctxBg, http.MethodPost, "url", getHttpBody("{}"))
				expectErrToBeNil(err)
				var output json.RawMessage
				// -- WHEN --
				err = DecodeRequestBody(req, &output)
				// -- THEN --
				expectErrToBeNil(err)
			})
		})

		When("body is nil", func() {
			It("returns error", func() {
				// -- GIVEN --
				req, err := http.NewRequestWithContext(ctxBg, http.MethodPost, "url", nil)
				expectErrToBeNil(err)
				var output json.RawMessage
				// -- WHEN --
				err = DecodeRequestBody(req, &output)
				// -- THEN --
				expectErrToNotBeNil(err)
			})
		})

		When("io.ReadAll error", func() {
			It("returns error", func() {
				// -- GIVEN --
				req, err := http.NewRequestWithContext(ctxBg, http.MethodPost, "url", errReadCloser(1))
				expectErrToBeNil(err)
				var output json.RawMessage
				// -- WHEN --
				err = DecodeRequestBody(req, &output)
				// -- THEN --
				expectErrToNotBeNil(err)
			})
		})

		When("json unmarshal error", func() {
			It("returns error", func() {
				// -- GIVEN --
				req, err := http.NewRequestWithContext(ctxBg, http.MethodPost, "url", getHttpBody("{]"))
				expectErrToBeNil(err)
				var output json.RawMessage
				// -- WHEN --
				err = DecodeRequestBody(req, &output)
				// -- THEN --
				expectErrToNotBeNil(err)
			})
		})
	})

	Describe("SafelyCloseBody tests", func() {
		When("body is not nil", func() {
			It("successfully closed", func() {
				// -- GIVEN --
				body := getHttpBody("{}")
				req, err := http.NewRequestWithContext(ctxBg, http.MethodPost, "url", body)
				expectErrToBeNil(err)
				// -- WHEN --
				SafelyCloseBody(req.Body)
				// -- THEN --
				err = body.Close()
				expectErrToBeNil(err)
			})
		})

		When("body is nil", func() {
			It("does not panic", func() {
				// -- GIVEN --
				body := getHttpBody("{}")
				req, err := http.NewRequestWithContext(ctxBg, http.MethodPost, "url", body)
				expectErrToBeNil(err)
				// -- WHEN --
				// -- THEN --
				Expect(func() { SafelyCloseBody(req.Body) }).ToNot(Panic())
			})
		})
	})
})

// errReadCloser - Ref: https://stackoverflow.com/a/45126402/6323360
type errReadCloser int

func (m errReadCloser) Read(_ []byte) (n int, err error) {
	return 0, errors.New("some error")
}

func (m errReadCloser) Close() error {
	return nil
}

type mockClientStruct struct {
	mock.Mock
	numDoCount int
}

func (m *mockClientStruct) Do(req *http.Request) (*http.Response, error) {
	if m.numDoCount > 0 {
		m.numDoCount -= 1
		return nil, errors.New("some error")
	}

	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func expectErrToBeNil(errInput error) {
	Expect(errInput).To(BeNil())
}

func expectErrToNotBeNil(errInput error) {
	Expect(errInput).To(Not(BeNil()))
}

func getHttpBody(input string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(input)))
}

func resetMonkeyPatching() {
}
