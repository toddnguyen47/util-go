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

func GetFormattedString(epochMilli int64) string {
	return time.UnixMilli(epochMilli).In(time.UTC).Format(ISO8601)
}

// ParseAndGetEpoch - Parse a time string and return its epoch
func ParseAndGetEpoch(timeInput string) (int64, error) {
	time1, err := Parse(timeInput)
	if err != nil {
		var results int64
		return results, err
	}
	return GetEpoch(time1), nil
}
