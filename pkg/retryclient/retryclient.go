package retryclient

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"time"
)

type clientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

type RetryConfig struct {
	RetryTimes uint
	SleepTime  time.Duration
	Client     clientInterface
	Request    *http.Request
}

/*
Retry

Retry a client call with exponential backoff. Ref: https://docs.aws.amazon.com/general/latest/gr/api-retries.html

Example usage:

	resp, err := Retry(RetryConfig{
			RetryTimes: 3,
			SleepTime:  100 * time.Millisecond,
			Client:     client,
			Request:    req,
		})
*/
func Retry(retryConfig RetryConfig) (*http.Response, error) {
	retries := uint(0)
	retry := true

	resp := &http.Response{
		StatusCode: http.StatusInternalServerError,
		Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
	}
	// Default err to failure
	var err = errors.New("retry failure")

	// Reuse body if there is one
	var postPayload []byte
	if retryConfig.Request.Body != nil {
		var err2 error
		postPayload, err2 = io.ReadAll(retryConfig.Request.Body)
		if err2 != nil {
			return resp, err2
		}
	}

	for retry && retries < retryConfig.RetryTimes {
		if retries > 0 {
			sleepTime := (1 << retries) * retryConfig.SleepTime
			time.Sleep(sleepTime)
		}
		req := retryConfig.Request
		if postPayload != nil {
			req.Body = io.NopCloser(bytes.NewReader(postPayload))
		}
		resp, err = retryConfig.Client.Do(req)
		if err == nil && resp.StatusCode >= http.StatusOK && resp.StatusCode <= 299 {
			retry = true
		}

		retries += 1
	}

	return resp, err
}
