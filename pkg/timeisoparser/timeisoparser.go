package timeisoparser

import "time"

// TimeLayouts - We have to use a specific time, as defined here:
// https://pkg.go.dev/time#Layout
const (
	ISO8601          = "2006-01-02T15:04:05.000Z"
	ISO8601NoPeriods = "2006-01-02T15:04:05Z"
)

var timeLayoutList = []string{ISO8601, ISO8601NoPeriods}

// Parse through all possible formats
func Parse(timeInput string) (time.Time, error) {
	var err error
	var timeOutput time.Time
	for _, timeLayout := range timeLayoutList {
		timeOutput, err = time.Parse(timeLayout, timeInput)
		if err == nil {
			break
		}
	}
	return timeOutput, err
}

// GetEpoch - function to consistently convert time.Time to int64
func GetEpoch(timeInput time.Time) int64 {
	return timeInput.UnixMilli()
}
