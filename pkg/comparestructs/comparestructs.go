package comparestructs

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

// Compare - Compare two fields with optional options. Return an empty string if two structs are the same,
// otherwise return a string to denote what is different.
// To ignore a field, please see `comparestructs_test.go`
func Compare(x, y interface{}, opts ...cmp.Option) string {
	diff := cmp.Diff(x, y, opts...)
	if diff == "" {
		return ""
	} else {
		return fmt.Sprintf("The two items are NOT equal. (-want +got):\n%s", diff)
	}
}
