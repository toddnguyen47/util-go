package safederef

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type EventTest struct {
	Id     *string
	IdList *[]string
}

func Test_GivenNilPtr_WhenDeref_ThenReturnEmptyString(t *testing.T) {
	event := EventTest{}

	actual := Deref(event.Id)

	var expected string
	assert.Equal(t, expected, actual)
}

func Test_GivenProperString_WhenDeref_ThenReturnThatString(t *testing.T) {
	id := "Some ID"
	event := EventTest{
		Id: &id,
	}

	actual := Deref(event.Id)

	assert.Equal(t, id, actual)
}

func Test_GivenNilPtr_WhenDerefStringSlice_ThenReturnEmptyString(t *testing.T) {
	event := EventTest{}

	actual := DerefSlice(event.IdList)

	expected := make([]string, 0)
	assert.Equal(t, expected, actual)
}

func Test_GivenProperString_WhenDerefStringSlice_ThenReturnThatString(t *testing.T) {
	ids := []string{"Some ID", "Some ID2", "Some ID3", "Some ID3", "Some ID4"}
	event := EventTest{
		IdList: &ids,
	}

	actual := DerefSlice(event.IdList)

	assert.Equal(t, ids, actual)
}
