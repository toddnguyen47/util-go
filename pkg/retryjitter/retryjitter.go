package retryjitter

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

var _reader = rand.Reader
var _minSleepTimeMillis int64 = 50

// Retry - retry with exponential backoff and jitter. Default timeout is 100 milliseconds for the first sleep.
// If you want to customize the sleep time, call RetryWithTimeout().
//
// Ref: https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/
//
// # Sample Usage
//
//	func Example(ctx context.Context) {
//		var someValue int
//		err := retryjitter.Retry(ctx, 3, func() error {
//			var innerErr error
//			someValue, innerErr = doSomeWork(ctx)
//			return innerErr
//		})
//	}
func Retry(ctx context.Context, retryTimes int, funcToRetry func() error) error {
	return RetryWithTimeout(ctx, retryTimes, 100, funcToRetry)
}

// RetryWithTimeout - Same as Retry, except passing in a timeout. We have to pass an int as we need to
// randomize the time we are sleeping, from 0 to maxTime (maxTime is calculated per retry count).
func RetryWithTimeout(_ context.Context, retryTimes int, timeoutMilliseconds int, funcToRetry func() error) error {
	count := 0
	keepRetrying := true
	var err error

	timeoutMillisInner := int64(timeoutMilliseconds)
	if timeoutMillisInner <= 0 {
		fmt.Println("timeout passed is less than or equal to zero; defaulting to 100")
		timeoutMillisInner = 100
	}

	for ; count <= retryTimes && keepRetrying; count += 1 {
		if count > 0 {
			maxSleep := timeoutMillisInner << (count - 1)
			maxRange := big.NewInt(maxSleep + 1 - _minSleepTimeMillis)
			sleepTime, err2 := rand.Int(_reader, maxRange)
			if err2 != nil {
				sleepTime = maxRange
			}
			sleepTimeInt64 := sleepTime.Int64() + _minSleepTimeMillis
			fmt.Printf("Current Count: %d, Previous Error: %s, Sleeping for: %d milliseconds\n",
				count, err.Error(), sleepTimeInt64)
			time.Sleep(time.Duration(sleepTimeInt64) * time.Millisecond)
		}
		err = funcToRetry()
		if err == nil {
			keepRetrying = false
		}
	}

	return err
}
