package timeisoparser

import (
	"strings"
	"time"
)

// TimeLayouts - We have to use a specific time, as defined here:
// https://pkg.go.dev/time#Layout
//
// The reason we use `.000` is because want to include milliseconds always.
//
// The documentation states:
// A comma or decimal point followed by one or more zeros represents a fractional second, printed to the given number of decimal places. A comma or decimal point followed by one or more nines represents a fractional second, printed to the given number of decimal places, with trailing zeros removed.
const (
	ISO8601                 string = "2006-01-02T15:04:05Z"
	ISO8601Millis           string = "2006-01-02T15:04:05.000Z"
	ISO8601DateOnly         string = "2006-01-02"
	ISO8601FileName         string = "2006-01-02T15-04-05Z"
	ISO8601FileNameDateOnly string = "2006-01-02"

	GoUnixEpoch = int64(1136239445)
)

var timeLayoutList = []string{ISO8601, ISO8601Millis}

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

func GetFormattedISO8601MillisString(epochMilli int64) string {
	return time.UnixMilli(epochMilli).In(time.UTC).Format(ISO8601Millis)
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

// GetDatesInRangeStr - https://stackoverflow.com/a/58480030/6323360
func GetDatesInRangeStr(rangeStart, rangeEnd string) []string {
	rangeStartTime, err := Parse(rangeStart)
	if err != nil {
		return []string{}
	}
	rangeEndTime, err := Parse(rangeEnd)
	if err != nil {
		return []string{}
	}

	return GetDatesInRange(rangeStartTime, rangeEndTime)
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
