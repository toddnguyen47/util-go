package jsonutils

import (
	"bytes"
	"encoding/json"
)

// MarshalNoEscapeHtml - marshal JSON without escaping HTML
// Ref: https://stackoverflow.com/a/28596225/6323360
// Ref for trimming `\n`: https://stackoverflow.com/questions/28595664/how-to-stop-json-marshal-from-escaping-and#comment122847570_28596225
func MarshalNoEscapeHtml(v any) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(v)
	if err != nil {
		return make([]byte, 0), err
	}
	b1 := bytes.TrimRight(buf.Bytes(), "\n")
	return b1, nil
}
