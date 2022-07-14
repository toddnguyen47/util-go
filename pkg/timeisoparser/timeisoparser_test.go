package timeisoparser

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Start MyIsoTime `json:"start"`
}

func Test_GivenValidTimeIso8601_WhenUnmarshal_ThenReturnParsedTime(t *testing.T) {
	s1 := `{"start": "2022-02-01T00:00:00.000Z"}`
	var ts1 testStruct
	err := json.Unmarshal([]byte(s1), &ts1)

	assert.Nil(t, err)
	expectedEpoch := int64(1643673600000)
	assert.Equal(t, expectedEpoch, GetEpoch(ts1.Start.GetTime()))
}

func Test_GivenValidTimeIso8601NoMillis_WhenUnmarshal_ThenReturnParsedTime(t *testing.T) {
	s1 := `{"start": "2022-02-01T00:00:00Z"}`
	var ts1 testStruct
	err := json.Unmarshal([]byte(s1), &ts1)

	assert.Nil(t, err)
	expectedEpoch := int64(1643673600000)
	assert.Equal(t, expectedEpoch, GetEpoch(ts1.Start.GetTime()))
}

func Test_GivenInvalidTimeIso8601_WhenUnmarshal_ThenReturnErr(t *testing.T) {
	s1 := `{"start": "asdf"}`
	var ts1 testStruct

	err := json.Unmarshal([]byte(s1), &ts1)

	var expectedTime time.Time
	assert.NotNil(t, err)
	assert.Equal(t, GetEpoch(expectedTime), GetEpoch(ts1.Start.GetTime()))
}
