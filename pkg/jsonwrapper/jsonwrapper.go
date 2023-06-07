package jsonwrapper

import (
	"bytes"
	"encoding/json"
	"io"
)

type Interface interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

type defaultJsonWrapper struct{}

func NewDefaultJsonWrapper() Interface {
	return new(defaultJsonWrapper)
}

func (d defaultJsonWrapper) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (d defaultJsonWrapper) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func MarshalNoEscapeHtml(v any) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(v)
	if err != nil {
		return make([]byte, 0), err
	}
	return io.ReadAll(&buf)
}
