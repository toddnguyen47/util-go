package concurrentlist

import "sync"

// ConcurrentList - simple list that aims to be thread safe
type ConcurrentList[T any] interface {
	GetList() []T

	Append(val T)
	Get(index int) (T, bool)

	// Delete - this call can be very slow as the underlying slice will have to shift
	Delete(index int) (T, bool)

	// Update - returns false if update fails
	Update(index int, val T) bool

	IsEmpty() bool
	Size() int
}

func NewConcurrentList[T any]() ConcurrentList[T] {
	return &impl[T]{
		count:          0,
		underlyingList: make([]T, 0),
	}
}

type impl[T any] struct {
	mutex          sync.RWMutex
	count          int
	underlyingList []T
}

func (i1 *impl[T]) GetList() []T {
	i1.mutex.RLock()
	defer i1.mutex.RUnlock()
	newList := make([]T, i1.count)
	copy(newList, i1.underlyingList)
	return newList
}

func (i1 *impl[T]) Get(index int) (T, bool) {
	i1.mutex.RLock()
	defer i1.mutex.RUnlock()
	var val T
	if !i1.isWithinRange(index) {
		return val, false
	}
	return i1.underlyingList[index], true
}

func (i1 *impl[T]) Append(val T) {
	i1.mutex.Lock()
	defer i1.mutex.Unlock()
	i1.count += 1
	i1.underlyingList = append(i1.underlyingList, val)
}

func (i1 *impl[T]) Update(index int, val T) bool {
	i1.mutex.Lock()
	defer i1.mutex.Unlock()
	if !i1.isWithinRange(index) {
		return false
	}
	i1.underlyingList[index] = val
	return true
}

func (i1 *impl[T]) Delete(index int) (T, bool) {
	i1.mutex.Lock()
	defer i1.mutex.Unlock()
	var val T
	if !i1.isWithinRange(index) {
		return val, false
	}
	val = i1.underlyingList[index]
	i1.underlyingList = append(i1.underlyingList[0:index], i1.underlyingList[index+1:]...)
	i1.count -= 1
	return val, true
}

func (i1 *impl[T]) IsEmpty() bool {
	i1.mutex.RLock()
	defer i1.mutex.RUnlock()
	val := i1.count == 0
	return val
}

func (i1 *impl[T]) Size() int {
	return i1.count
}

func (i1 *impl[T]) isWithinRange(index int) bool {
	return index >= 0 && index < i1.count
}
