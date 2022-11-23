package urldecodeutils

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// httpClient is used to pass client.
type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// SendRequestDecodeResponse - for output, pass in a pointer to a struct
// This function returns a status code and an error
func SendRequestDecodeResponse(ctx context.Context, client httpClient, req *http.Request,
	output interface{}) (int, error) {

	resp, err := client.Do(req)
	if err != nil {
		return http.StatusBadRequest, err
	}

	// If response is not in the 200 - 299 range
	if resp.StatusCode < http.StatusOK || resp.StatusCode > 299 {
		return resp.StatusCode, errors.New("status code is not in the 200 (OK) range")
	}

	defer SafelyCloseBody(resp.Body)

	b1, err := io.ReadAll(resp.Body)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = json.Unmarshal(b1, output)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return resp.StatusCode, nil
}

// DecodeRequestBody - for output, pass in a pointer to a struct
func DecodeRequestBody(r *http.Request, output interface{}) error {

	if r.Body == nil {
		return errors.New("post payload is nil")
	}

	// Get required parameters
	b1, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b1, output)
	if err != nil {
		return err
	}

	return nil
}

// SafelyCloseBody - you can also use a `defer` function for this
func SafelyCloseBody(body io.ReadCloser) {
	if body != nil {
		_ = body.Close()
	}
}
