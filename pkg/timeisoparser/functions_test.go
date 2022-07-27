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
	str1 := GetFormattedString(1643702820000)

	assert.Equal(t, "2022-02-01T08:07:00.000Z", str1)
}
