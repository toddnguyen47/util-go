package jsonwrapper

import "encoding/json"

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
