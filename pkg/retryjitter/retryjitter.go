package retryjitter

import (
	"time"

	"github.com/rs/zerolog/log"
	"github.com/toddnguyen47/util-go/pkg/randomutils"
)

var _minSleepTimeMillis int64 = 100
var _maxSleepTimeMillis int64 = 20 * 10_000 // 20 seconds

// Retry - retry with exponential backoff and jitter. Default timeout is 100 milliseconds for the first sleep.
// If you want to customize the sleep time, call RetryWithTimeout().
//
// Ref: https://docs.aws.amazon.com/sdkref/latest/guide/feature-retry-behavior.html#ExponentialBackoff
// Set max time to 20 seconds
//
// Ref: https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/
//
// # Sample Usage
//
//	func Example() {
//		var someValue int
//		err := retryjitter.Retry(3, func() error {
//			var innerErr error
//			someValue, innerErr = doSomeWork()
//			return innerErr
//		})
//	}
func Retry(retryTimes int, funcToRetry func() error) error {
	return RetryWithTimeout(retryTimes, 100, funcToRetry)
}

// RetryWithTimeout - Same as Retry, except passing in a timeout. We have to pass an int as we need to
// randomize the time we are sleeping, from 0 to maxTime (maxTime is calculated per retry count).
func RetryWithTimeout(retryTimes int, timeoutMilliseconds int64, funcToRetry func() error) error {
	count := 0
	keepRetrying := true
	var err error

	for ; count <= retryTimes && keepRetrying; count++ {
		SleepIncrementalBackoffJitter(count, timeoutMilliseconds)
		err = funcToRetry()
		if err == nil {
			keepRetrying = false
		}
	}

	return err
}

// SleepIncrementalBackoffJitter - Sleep with incremental backoff and jitter.
// Reference: https://docs.aws.amazon.com/general/latest/gr/api-retries.html
// (1 << n) is equivalent to (2^n). Max sleep time will be 20 seconds (20_000 milliseconds)
func SleepIncrementalBackoffJitter(count int, sleepTimeMillis int64) {
	if count > 0 {
		timeoutMillisInner := int64(sleepTimeMillis)
		if timeoutMillisInner <= 0 {
			log.Info().Msg("timeout passed is less than or equal to zero; defaulting to 100 milliseconds")
			// Defaults to 100 milliseconds
			timeoutMillisInner = 100
		}
		maxSleep := timeoutMillisInner << (count - 1)
		sleepTimeInt64 := randomutils.GetRandomWithMin(_minSleepTimeMillis, maxSleep)
		if sleepTimeInt64 > _maxSleepTimeMillis {
			sleepTimeInt64 = _maxSleepTimeMillis
		}
		log.Info().
			Int("count", count).
			Int64("sleeping for x milliseconds", sleepTimeInt64).
			Msg("SleepIncrementalBackoffJitter logging")
		time.Sleep(time.Duration(sleepTimeInt64) * time.Millisecond)
	}
}
