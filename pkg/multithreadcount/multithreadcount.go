package multithreadcount

import "sync"

func NewMultiThreadCount() MultiThreadCount {
	return MultiThreadCount{
		SuccessCount: new(uint32),
		ErrCount:     new(uint32),
		mutex:        new(sync.RWMutex),
		err:          nil,
	}
}

// MultiThreadCount
// mutex Ref: https://stackoverflow.com/a/52882045/6323360
type MultiThreadCount struct {
	SuccessCount *uint32
	ErrCount     *uint32
	mutex        *sync.RWMutex
	err          error
}

// GetError - get an error with a defer to unlock, so multiple goroutine can read it
func (e *MultiThreadCount) GetError() error {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.err
}

// SetError - set an error.
func (e *MultiThreadCount) SetError(val error) {
	e.mutex.Lock()
	e.err = val
	e.mutex.Unlock()
}
