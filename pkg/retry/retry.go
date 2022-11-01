package retry

import (
	"errors"
	"time"

	"github.com/rs/zerolog"
)

type GenericFunc func(arguments ...interface{}) error

// IncrementalRetry - Ref: https://docs.aws.amazon.com/general/latest/gr/api-retries.html
func IncrementalRetry(logger *zerolog.Logger, retryTimes int, sleepTime time.Duration,
	fn GenericFunc, arguments ...interface{}) error { // NOSONAR

	count := 0
	doRetry := true
	for ; doRetry && count < retryTimes; count++ {
		sleepAndLog(logger, count, sleepTime)

		err := fn(arguments...)
		if err == nil {
			doRetry = false
		}
	}

	if doRetry || count >= retryTimes {
		return errors.New("`fn` function was not called successfully")
	}

	return nil
}

func sleepAndLog(logger *zerolog.Logger, count int, sleepTime time.Duration) {

	if count > 0 {
		c1 := count - 1
		// 2^n == 1 << n
		s1 := time.Duration(1<<c1) * sleepTime
		logger.Info().Int64("sleeping for n milliseconds", s1.Milliseconds()).Int("count", count)
		time.Sleep(s1)
	}
}
