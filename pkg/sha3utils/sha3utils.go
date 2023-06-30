package sha3utils

import (
	"fmt"

	"golang.org/x/crypto/sha3"
)

type DigestJsonMarshal interface {
	CustomMarshal() ([]byte, error)
}

// ComputeHash - the input will have to implement the CustomMarshal() function. You can simply return
// `json.Marshal()`. This function will default hash length to 64.
// Ref: https://pkg.go.dev/golang.org/x/crypto/sha3
func ComputeHash(input DigestJsonMarshal) string {
	// Should be no error marshalling
	buf, _ := input.CustomMarshal()
	return ComputeHashBytes(buf)
}

// ComputeHashWithHashLen - the input will have to implement the CustomMarshal() function. You can simply return
// `json.Marshal()`
// Ref: https://pkg.go.dev/golang.org/x/crypto/sha3
func ComputeHashWithHashLen(input DigestJsonMarshal, hashLen int) string {
	// Should be no error marshalling
	buf, _ := input.CustomMarshal()
	return ComputeHashBytesWithHashLen(buf, hashLen)
}

// ComputeHashBytes - This function will default hash length to 64.
func ComputeHashBytes(buf []byte) string {
	// A hash needs to be 64 bytes long to have 256-bit collision resistance.
	return ComputeHashBytesWithHashLen(buf, 64)
}

func ComputeHashBytesWithHashLen(buf []byte, hashLen int) string {
	h := make([]byte, hashLen)
	// Compute a 64-byte hash of buf and put it in h.
	sha3.ShakeSum256(h, buf)
	return fmt.Sprintf("%x", h)
}
