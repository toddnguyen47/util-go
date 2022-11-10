package sha3utils

import (
	"fmt"

	"golang.org/x/crypto/sha3"
)

type DigestJsonMarshal interface {
	CustomMarshal() ([]byte, error)
}

// ComputeHash - the input will have to implement the CustomMarshal() function
// Ref: https://pkg.go.dev/golang.org/x/crypto/sha3
func ComputeHash(input DigestJsonMarshal) string {
	// Should be no error marshalling
	buf, _ := input.CustomMarshal()
	// A hash needs to be 64 bytes long to have 256-bit collision resistance.
	h := make([]byte, 64)
	// Compute a 64-byte hash of buf and put it in h.
	sha3.ShakeSum256(h, buf)
	return fmt.Sprintf("%x", h)
}
