package safederef

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type EventTest struct {
	Id *string
}

func Test_GivenNilPtr_WhenDeref_ThenReturnEmptyString(t *testing.T) {
	event := EventTest{}

	actual := DerefStr(event.Id)

	var expected string
	assert.Equal(t, expected, actual)
}

func Test_GivenProperString_WhenDeref_ThenReturnThatString(t *testing.T) {
	id := "Some ID"
	event := EventTest{
		Id: &id,
	}

	actual := DerefStr(event.Id)

	assert.Equal(t, id, actual)
}
