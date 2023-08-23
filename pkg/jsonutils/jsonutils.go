package jsonutils

import (
	"bytes"
	"encoding/json"
	"io"
)

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
