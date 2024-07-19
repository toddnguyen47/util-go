package mapequalfold

import (
	"fmt"
	"strings"
)

type MapEqualFold[T fmt.Stringer] map[string]T

func NewMapEqualFold[T fmt.Stringer]() MapEqualFold[T] {
	m := make(MapEqualFold[T])
	return m
}

func (m *MapEqualFold[T]) Get(key string) (T, bool) {
	val, ok := (*m)[strings.ToUpper(key)]
	return val, ok
}

func (m *MapEqualFold[T]) Set(key string, value T) {
	(*m)[strings.ToUpper(key)] = value
}
