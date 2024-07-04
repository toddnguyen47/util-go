package jsonutils

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
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

// IterateJson - recursively iterate a JSON and store data into `currentMap`.
// Keys will be in the form of "key1.0.key2".
// Lists will have their indices as a key.
//
// Sample usage:
//
//	var inputMap map[string]interface{}
//	// a1 is any input / struct
//	err := json.Unmarshal(a1, &inputMap)
//	if err != nil {
//		panic(err)
//	}
//	currentKey := make([]string, 0)
//	map1 := make(map[string]interface{})
//	jsonutils.IterateJson(inputMap, currentKey, map1)
//	// Do stuff with map1 as needed
func IterateJson(jsonData interface{}, currentKey []string, currentMap map[string]interface{}) {
	if jsonData == nil {
		setData(jsonData, currentKey, currentMap)
		return
	}

	type1 := strings.ToLower(reflect.TypeOf(jsonData).String())
	if strings.EqualFold(type1, "map[string]interface {}") {
		map1 := jsonData.(map[string]interface{})
		for key, val := range map1 {
			currentKey = append(currentKey, key)
			IterateJson(val, currentKey, currentMap)
			// Pop off list now
			currentKey = currentKey[0 : len(currentKey)-1]
		}
	} else if strings.EqualFold(type1, "[]interface {}") {
		l1 := jsonData.([]interface{})
		for idx, val := range l1 {
			currentKey = append(currentKey, strconv.Itoa(idx))
			IterateJson(val, currentKey, currentMap)
			// Pop off list now
			currentKey = currentKey[0 : len(currentKey)-1]
		}
	} else {
		setData(jsonData, currentKey, currentMap)
	}
}

func setData(jsonData interface{}, currentKey []string, currentMap map[string]interface{}) {
	key := strings.Join(currentKey, ".")
	currentMap[key] = jsonData
}
