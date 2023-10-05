package retryclient

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/toddnguyen47/util-go/pkg/clientinterface"
	"github.com/toddnguyen47/util-go/pkg/retryjitter"
)

type RetryConfig struct {
	RetryTimes uint
	SleepTime  time.Duration
	Client     clientinterface.Client
	Request    *http.Request
}

func (r *RetryConfig) validate() error {
	if r.RetryTimes == 0 {
		return fmt.Errorf("retry times need to be greater than zero")
	}
	return nil
}

// Retry
//
// Retry a client call with exponential backoff. Ref: https://docs.aws.amazon.com/general/latest/gr/api-retries.html
//
// Example usage:
//
//	resp, err := Retry(RetryConfig{
//		RetryTimes: 3,
//		SleepTime:  100 * time.Millisecond,
//		Client:     client,
//		Request:    req,
//	})
func Retry(retryConfig RetryConfig) (*http.Response, error) {

	resp := &http.Response{
		StatusCode: http.StatusInternalServerError,
		Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
	}
	retryErr := retryConfig.validate()
	if retryErr != nil {
		return resp, retryErr
	}

	// Reuse body if there is one
	var postPayload []byte
	if retryConfig.Request.Body != nil {
		var err error
		postPayload, err = io.ReadAll(retryConfig.Request.Body)
		if err != nil {
			return resp, err
		}
	}

	retryErr = retryjitter.Retry(int(retryConfig.RetryTimes), func() error {
		var innerErr error
		req := retryConfig.Request
		if postPayload != nil {
			req.Body = io.NopCloser(bytes.NewReader(postPayload))
		}
		resp, innerErr = retryConfig.Client.Do(req)
		if innerErr != nil {
			return innerErr
		}
		statusCode := resp.StatusCode
		if statusCode < http.StatusOK || resp.StatusCode > 299 {
			newErr := fmt.Errorf("status code is NOT OK; status code: %d", statusCode)
			return newErr
		}
		return nil
	})

	return resp, retryErr
}
