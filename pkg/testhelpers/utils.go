package testhelpers

import (
	"os"
)

func ReadInTestFile(fileName string) []byte {
	b1, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return b1
}
