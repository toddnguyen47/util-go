package sha3utils

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Id       *int    `json:"id"`
	Name     *string `json:"name"`
	Position *string `json:"position"`
}

func (t testStruct) CustomMarshal() ([]byte, error) {
	return json.Marshal(t)
}

func Test_GivenTwoDifferentSourceInfo_When_ThenDigestShouldBeSame(t *testing.T) {
	// -- ARRANGE --
	var ts1 testStruct
	var ts2 testStruct
	var ts3 testStruct
	getTestStruct(t, &ts1)
	getTestStruct(t, &ts2)
	getTestStruct(t, &ts3)
	// -- ACT --
	digest1 := ComputeHash(ts1)
	digest2 := ComputeHash(ts2)
	digest3 := ComputeHash(ts3)
	// -- ASSERT --
	expectedDigest := "a3028ede34709cf8f497fd7d9501049243d7eb75bdcf037085dae333313cd2dc617c3eb11eff4cf2519be1a40e47b6ae8eca013ae7f8c923324217986b95db26"
	assert.Equal(t, digest1, digest2)
	assert.Equal(t, digest2, digest3)
	assert.Equal(t, expectedDigest, digest1)
	assert.Equal(t, expectedDigest, digest2)
	assert.Equal(t, expectedDigest, digest3)
}

func getTestStruct(_ *testing.T, t1 *testStruct) {
	*t1 = testStruct{
		Id:       makePtr(142),
		Name:     makePtr("name"),
		Position: makePtr("position"),
	}
}

func makePtr[T comparable](i T) *T {
	return &i
}
