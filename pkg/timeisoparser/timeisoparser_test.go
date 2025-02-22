package timeisoparser

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GivenIso8601FormatWithMillis_When_ThenNoErrIsReturned(t *testing.T) {
	timeStr := "2022-02-01T08:07:00.000Z"

	timeOutput, err := Parse(timeStr)

	assert.Nil(t, err)
	assert.Equal(t, time.February, timeOutput.Month())
	assert.Equal(t, 8, timeOutput.Hour())
	assert.Equal(t, 7, timeOutput.Minute())
}

func Test_GivenIso8601FormatNoMillis_When_ThenNoErrIsReturned(t *testing.T) {
	timeStr := "2022-02-01T08:07:00Z"

	timeOutput, err := Parse(timeStr)

	assert.Nil(t, err)
	assert.Equal(t, time.February, timeOutput.Month())
	assert.Equal(t, 8, timeOutput.Hour())
	assert.Equal(t, 7, timeOutput.Minute())
}

func Test_GivenIso8601FormatNoMillis_WhenParseAndGetEpoch_ThenNoErrIsReturned(t *testing.T) {
	timeStr := "2022-02-01T08:07:00Z"

	timeOutput, err := ParseAndGetEpoch(timeStr)

	assert.Nil(t, err)
	assert.Equal(t, int64(1643702820000), timeOutput)
}

func Test_GivenIncorrectFormat_When_ThenErrIsNotNil(t *testing.T) {
	timeStr := "2022-02-01T08:07:00"

	_, err := Parse(timeStr)

	assert.NotNil(t, err)
}

func Test_GivenIncorrectFormat_WhenParseAndGetEpoch_ThenErrIsNotNil(t *testing.T) {
	timeStr := "2022-02-01T08:07:00"

	_, err := ParseAndGetEpoch(timeStr)

	assert.NotNil(t, err)
}

func Test_GivenTime_WhenGetEpoch_ThenEpochIsReturned(t *testing.T) {
	timeStr := "2022-02-01T08:07:00.000Z"
	timeOutput, err := Parse(timeStr)
	assert.Nil(t, err)

	epoch := GetEpoch(timeOutput)

	assert.Equal(t, int64(1643702820000), epoch)
}

func Test_GivenTime_WhenGetTtl_ThenEpochIsReturned(t *testing.T) {
	timeStr := "2022-02-01T08:07:00.000Z"
	timeOutput, err := Parse(timeStr)
	assert.Nil(t, err)

	epoch := GetTimeToLive(timeOutput)

	assert.Equal(t, int64(1643702820), epoch)
}

func Test_GivenEpoch_WhenFormatting_ThenNoErrIsReturned(t *testing.T) {
	str1 := GetFormattedISO8601MillisString(1643702820000)

	assert.Equal(t, "2022-02-01T08:07:00.000Z", str1)
}

func Test_GivenWithinRangeStart_ThenReturnTrue(t *testing.T) {
	timeInput := "2022-02-01T08:07:00.000Z"
	start := "2022-02-01T08:07:00.000Z"
	end := "2022-02-01T10:07:00.000Z"

	withinRange := IsWithinRangeInclusive(timeInput, start, end)

	assert.True(t, withinRange)
}

func Test_GivenTimeInputParsingError_ThenReturnFalse(t *testing.T) {
	timeInput := "asdf"
	start := "2022-02-01T08:07:00.000Z"
	end := "2022-02-01T10:07:00.000Z"

	withinRange := IsWithinRangeInclusive(timeInput, start, end)

	assert.False(t, withinRange)
}

func Test_GivenStartParsingError_ThenReturnFalse(t *testing.T) {
	timeInput := "2022-02-01T08:07:00.000Z"
	start := "asdf"
	end := "2022-02-01T10:07:00.000Z"

	withinRange := IsWithinRangeInclusive(timeInput, start, end)

	assert.False(t, withinRange)
}

func Test_GivenEndEmpty_ThenReturnTrue(t *testing.T) {
	timeInput := "2022-02-01T08:07:00.000Z"
	start := "2022-02-01T08:07:00.000Z"
	end := ""

	withinRange := IsWithinRangeInclusive(timeInput, start, end)

	assert.True(t, withinRange)
}

func Test_GivenEndParsingError_ThenReturnFalse(t *testing.T) {
	timeInput := "2022-02-01T08:07:00.000Z"
	start := "2022-02-01T08:07:00.000Z"
	end := "asdf"
	_ = NowUTC()

	withinRange := IsWithinRangeInclusive(timeInput, start, end)

	assert.False(t, withinRange)
}

func Test_GivenStartEndInSameDay_ThenReturnOneDay(t *testing.T) {
	startStr := "2022-02-28T23:00:00.000Z"
	endStr := "2022-02-28T23:30:00.000Z"

	list1 := GetDatesInRangeStr(startStr, endStr)

	assert.Equal(t, 1, len(list1))
}

func Test_GivenStartEndDifferentDays_ThenReturnTwoDays(t *testing.T) {
	startStr := "2022-02-28T23:00:00.000Z"
	endStr := "2022-03-01T02:00:00.000Z"

	list1 := GetDatesInRangeStr(startStr, endStr)

	assert.Equal(t, 2, len(list1))
}

func Test_GivenStartAfterEnd_ThenReturnZeroDays(t *testing.T) {
	startStr := "2022-02-28T23:00:00.000Z"
	endStr := "2022-02-28T22:59:00.000Z"

	list1 := GetDatesInRangeStr(startStr, endStr)

	assert.Equal(t, 0, len(list1))
}

func Test_GivenStartParseError_ThenReturnZeroDays(t *testing.T) {
	startStr := "asdf"
	endStr := "2022-02-28T22:59:00.000Z"

	list1 := GetDatesInRangeStr(startStr, endStr)

	assert.Equal(t, 0, len(list1))
}

func Test_GivenEndParseError_ThenReturnZeroDays(t *testing.T) {
	startStr := "2022-02-28T23:00:00.000Z"
	endStr := "asdf"

	list1 := GetDatesInRangeStr(startStr, endStr)

	assert.Equal(t, 0, len(list1))
}

func Test_GivenMaxTime_ThenReturnMaxTime(t *testing.T) {
	maxTime := MaxTime()
	assert.True(t, time.Date(9999, 12, 31, 23, 59, 59, 999999999, time.UTC).Equal(maxTime))
}

func Test_GivenMaxTimeISOMillis_ThenReturnMaxTimeMillis(t *testing.T) {
	maxTime := MaxTimeISOMillis()
	assert.Equal(t, "9999-12-31T23:59:59.999Z", maxTime)
}
