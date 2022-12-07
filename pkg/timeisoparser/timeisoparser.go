package timeisoparser

import (
	"time"
)

// TimeLayouts - We have to use a specific time, as defined here:
// https://pkg.go.dev/time#Layout
const (
	ISO8601         = "2006-01-02T15:04:05Z"
	ISO8601Millis   = "2006-01-02T15:04:05.999Z"
	ISO8601DateOnly = "2006-01-02"
)

var timeLayoutList = []string{ISO8601, ISO8601Millis}

// MyIsoTime - Ref: https://stackoverflow.com/a/39180230/6323360
type MyIsoTime struct {
	// Embed the `time.Time` struct
	time time.Time
}

// UnmarshalJSON - Implementing `Unmarshaler` interface
func (m *MyIsoTime) UnmarshalJSON(bytes1 []byte) error {
	str1 := string(bytes1)

	// Remove beginning and end quotes
	str1 = str1[1 : len(str1)-1]

	time1, err := Parse(str1)
	if err != nil {
		return err
	}
	m.time = time1
	return nil
}

func (m *MyIsoTime) GetTime() time.Time {
	return m.time
}
