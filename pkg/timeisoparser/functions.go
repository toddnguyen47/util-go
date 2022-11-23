package timeisoparser

import (
	"strings"
	"time"
)

func NowUTC() time.Time {
	return time.Now().UTC()
}

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

// GetTimeToLive - Amazon AWS DynamoDB TTL seems to only accept unix seconds, not unix milliseconds
func GetTimeToLive(timeInput time.Time) int64 {
	return timeInput.Unix()
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

// IsWithinRangeInclusive - Find if `timeInput` is within range, e.g. start <= timeInput <= end
func IsWithinRangeInclusive(timeInput, start, end string) bool {
	timeInputEpoch, err := ParseAndGetEpoch(timeInput)
	if err != nil {
		return false
	}

	startEpoch, err := ParseAndGetEpoch(start)
	// Defaults from start to false
	isGeqStart := false
	if err == nil {
		isGeqStart = startEpoch <= timeInputEpoch
	}

	isLeqEnd := false
	if strings.TrimSpace(end) == "" {
		isLeqEnd = true
	} else {
		endEpoch, err2 := ParseAndGetEpoch(end)
		if err2 == nil {
			isLeqEnd = timeInputEpoch <= endEpoch
		}
	}

	return isGeqStart && isLeqEnd
}

// GetDatesInRange - https://stackoverflow.com/a/58480030/6323360
func GetDatesInRange(rangeStart, rangeEnd time.Time) []string {
	times := make([]string, 0)
	if !rangeStart.Before(rangeEnd) {
		return times
	}

	start := time.Date(rangeStart.Year(), rangeStart.Month(), rangeStart.Day(), 0, 0, 0, 0,
		rangeStart.Location())
	for ; !start.After(rangeEnd); start = start.AddDate(0, 0, 1) {
		s1 := start.Format(ISO8601DateOnly)
		times = append(times, s1)
	}

	return times
}
