package jsonwrapper

import "encoding/json"

func NewJsonWrapper() Interface {
	return &impl{}
}

type impl struct{}

func (i2 *impl) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (i2 *impl) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
