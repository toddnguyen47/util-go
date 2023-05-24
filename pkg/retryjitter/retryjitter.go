package retryjitter

import (
	"crypto/rand"
	"math/big"
	"time"
)

var _reader = rand.Reader

// Retry - retry with exponential backoff and jitter.
//
// Ref: https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/
func Retry(retryTimes int, funcToRetry func() error) error {

	count := 0
	keepRetrying := true
	var err error

	for ; count <= retryTimes && keepRetrying; count += 1 {
		if count > 0 {
			maxSleep := big.NewInt(100 << (count - 1))
			sleepTime, err2 := rand.Int(_reader, maxSleep)
			if err2 != nil {
				sleepTime = maxSleep
			}
			sleepTimeInt64 := sleepTime.Int64()
			time.Sleep(time.Duration(sleepTimeInt64) * time.Millisecond)
		}
		err = funcToRetry()
		// TODO: If you wish, add logging of current count and current error
		if err == nil {
			keepRetrying = false
		}
	}

	return err
}
