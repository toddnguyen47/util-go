package utilhttp

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/toddnguyen47/util-go/sleep"
)

const ErrorHttpNotOk = "HTTP Request Status Code is not OK"

type QueryParam struct {
	Key   string
	Value string
}

func (q QueryParam) String() string {
	return fmt.Sprintf("%s:%s", q.Key, q.Value)
}

// RequestGet - Utility method to call a GET request while also checking for OK status after a call
func RequestGet(url string) ([]byte, error) {
	fmt.Printf("GET URL: '%s'\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(ErrorHttpNotOk)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

// RequestGetWithQueryParams - Utility method to call a GET request while also checking for OK status after a call
func RequestGetWithQueryParams(url string, queryParams *[]QueryParam) ([]byte, error) {
	if queryParams == nil {
		panic("queryParams cannot be nil")
	}

	url = url + "?"
	for _, queryParam := range *queryParams {
		url = url + queryParam.Key + "=" + queryParam.Value + "&"
		fmt.Println(queryParam)
	}
	// Remove trailing &
	url = url[:len(url)-1]

	return RequestGet(url)
}

// RequestPut - Utility method to call a PUT request while also checking for OK status after a call
func RequestPut(url string) error {
	var payloadBytes io.Reader
	return requestMethodWrapper(http.MethodPut, url, payloadBytes)
}

// RequestPutWithPayload - Utility method to call a PUT request while also checking for OK status after a call
func RequestPutWithPayload(url string, payload *[]byte) error {
	if payload == nil {
		log.Fatal("Payload cannot be nil!")
	}
	payloadBytes := bytes.NewBuffer(*payload)
	return requestMethodWrapper(http.MethodPut, url, payloadBytes)
}

// RequestDelete - Utility method to call a DELETE request while also checking for OK status after a call
func RequestDelete(url string) error {
	var payloadBytes io.Reader
	return requestMethodWrapper(http.MethodDelete, url, payloadBytes)
}

// RequestDeleteWithPayload - Utility method to call a DELETE request while also checking for OK status after a call
func RequestDeleteWithPayload(url string, payload *[]byte) error {
	if payload == nil {
		log.Fatal("Payload cannot be nil!")
	}
	payloadBytes := bytes.NewBuffer(*payload)
	return requestMethodWrapper(http.MethodDelete, url, payloadBytes)
}

// ********************************************************
// PRIVATE FUNCTIONS
// ********************************************************

func requestMethodWrapper(method string, url string, payloadBytes io.Reader) error {
	// Set up request
	request, err := http.NewRequest(method, url, payloadBytes)
	if err != nil {
		return err
	}
	request.Header.Set("Content-type", "application/json")

	// Call request
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return errors.New(ErrorHttpNotOk)
	}
	fmt.Printf("Finished %s request\n", http.MethodPut)
	sleep.SleepsAndLog(2)
	return nil
}
