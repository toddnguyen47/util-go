package retryjitter

import (
	"math/rand"
	"time"
)

// Retry - retry with exponential backoff and jitter.
//
// Ref: https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/
func Retry(retryTimes int, funcToRetry func() error) error {

	now := time.Now()
	myRand := rand.New(rand.NewSource(now.UnixMilli()))
	count := 0
	keepRetrying := true
	var err error

	for ; count < retryTimes && keepRetrying; count += 1 {
		if count > 0 {
			maxSleep := 100 << (count - 1)
			sleepTime := myRand.Intn(maxSleep)
			time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		}
		err = funcToRetry()
		if err == nil {
			keepRetrying = false
		}
	}

	return err
}
