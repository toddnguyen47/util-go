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
	s1 := `{"start":"2022-02-01T00:00:00.001Z"}`
	var ts1 testStruct
	err := json.Unmarshal([]byte(s1), &ts1)

	assert.Nil(t, err)
	expectedEpoch := int64(1643673600001)
	assert.Equal(t, expectedEpoch, GetEpoch(ts1.Start.Time()))
	b1, err := json.Marshal(ts1)
	assert.Nil(t, err)
	assert.Equal(t, s1, string(b1))
	assert.Equal(t, int64(1136239445), GoUnixEpoch)
}

func Test_GivenValidTimeIso8601NoMillis_WhenUnmarshal_ThenReturnParsedTime(t *testing.T) {
	s1 := `{"start":"2022-02-01T00:00:00Z"}`
	var ts1 testStruct
	err := json.Unmarshal([]byte(s1), &ts1)

	assert.Nil(t, err)
	expectedEpoch := int64(1643673600000)
	assert.Equal(t, expectedEpoch, GetEpoch(ts1.Start.Time()))
	b1, err := json.Marshal(ts1)
	assert.Nil(t, err)
	expectedS1Millis := `{"start":"2022-02-01T00:00:00.000Z"}`
	assert.Equal(t, expectedS1Millis, string(b1))
}

func Test_GivenInvalidTimeIso8601_WhenUnmarshal_ThenReturnErr(t *testing.T) {
	s1 := `{"start":"asdf"}`
	var ts1 testStruct

	err := json.Unmarshal([]byte(s1), &ts1)

	var expectedTime time.Time
	assert.NotNil(t, err)
	assert.Equal(t, GetEpoch(expectedTime), GetEpoch(ts1.Start.Time()))
}
