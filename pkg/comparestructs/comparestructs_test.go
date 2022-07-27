package comparestructs

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	IntValue        int
	StrValue        string
	IgnoreThisField time.Time
}

func Test_GivenTwoStructsSame_WhenComparing_ThenDiffIsEmptyString(t *testing.T) {
	testStruct1 := testStruct{
		IntValue:        42,
		StrValue:        "42",
		IgnoreThisField: time.Now(),
	}
	testStruct2 := testStruct{
		IntValue:        42,
		StrValue:        "42",
		IgnoreThisField: time.Now(),
	}

	diff := Compare(testStruct1, testStruct2, cmpopts.IgnoreFields(testStruct{}, "IgnoreThisField"))

	assert.Equal(t, "", diff)
}

func Test_GivenTwoStructsDiff_WhenComparing_ThenDiffIsErrorString(t *testing.T) {
	testStruct1 := testStruct{
		IntValue:        42,
		StrValue:        "42",
		IgnoreThisField: time.Now(),
	}
	testStruct2 := testStruct{
		IntValue:        50,
		StrValue:        "50",
		IgnoreThisField: time.Now(),
	}

	diff := Compare(testStruct1, testStruct2, cmpopts.IgnoreFields(testStruct{}, "IgnoreThisField"))

	fmt.Println(diff)
	assert.NotEqual(t, "", diff)
}
