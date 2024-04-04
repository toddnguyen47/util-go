package timeisoparser

import (
	"fmt"
	"time"
)

// MyIsoTime - Ref: https://stackoverflow.com/a/39180230/6323360
// Ignore any error about pointer value and receiver, as stdlib also uses mixed receivers
// exclusively for JSON marshal:
// https://pkg.go.dev/encoding/json@go1.19#RawMessage.MarshalJSON
//
// This type is used to marshal / unmarshal JSON. To convert from `time.Time`,
// cast it with MyIsoTime(value).
//
// To convert back to `time.Time`, use MyIsoTime.Time()
type MyIsoTime time.Time

// UnmarshalJSON - Implementing `Unmarshaler` interface
func (m *MyIsoTime) UnmarshalJSON(bytes1 []byte) error {
	str1 := string(bytes1)

	// Remove beginning and end quotes
	str1 = str1[1 : len(str1)-1]

	time1, err := Parse(str1)
	if err != nil {
		return err
	}
	*m = MyIsoTime(time1)
	return nil
}

func (m MyIsoTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", m.Time().Format(ISO8601Millis))
	return []byte(stamp), nil
}

func (m MyIsoTime) Time() time.Time {
	return time.Time(m)
}
