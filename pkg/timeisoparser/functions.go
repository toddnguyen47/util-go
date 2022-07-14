package timeisoparser

import "time"

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
