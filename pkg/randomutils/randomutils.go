package randomutils

import (
	"crypto/rand"
	"math/big"
)

// Monkey patching for tests
var (
	_reader = rand.Reader
)

func GetRandomWithMin(minNumber int64, maxNumber int64) int64 {
	if maxNumber < minNumber {
		minNumber, maxNumber = maxNumber, minNumber
	}
	if maxNumber <= 0 {
		maxNumber = 0
	}
	maxRangeInt64 := maxNumber - minNumber
	// We need to add 1 as rand.Int() gets random number from [0, max); e.g. max is exclusive.
	randomNumberBigInt, err := rand.Int(_reader, big.NewInt(maxRangeInt64+1))
	if err != nil {
		randomNumberBigInt = big.NewInt(maxRangeInt64)
	}
	randomNumber := randomNumberBigInt.Int64()
	return randomNumber + minNumber
}
