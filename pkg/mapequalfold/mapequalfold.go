package mapequalfold

import (
	"fmt"
	"strings"
)

// MapEqualFoldIface - interface that defines a map with case-insensitive keys.
// T is here is defined as the type of the value in the map.
type MapEqualFoldIface[T fmt.Stringer] interface {
	Get(key string) (T, bool)
	Set(key string, value T)
}

// MapEmptyStruct - a set-like example with an empty struct.
type MapEmptyStruct map[string]*EmptyStruct

func NewMapEmptyStruct() MapEqualFoldIface[*EmptyStruct] {
	m := make(MapEmptyStruct)
	return &m
}

func (m *MapEmptyStruct) Get(key string) (*EmptyStruct, bool) {
	val, ok := (*m)[strings.ToUpper(key)]
	return val, ok
}

func (m *MapEmptyStruct) Set(key string, value *EmptyStruct) {
	(*m)[strings.ToUpper(key)] = value
}

type EmptyStruct struct{}

func (e *EmptyStruct) String() string {
	return ""
}
