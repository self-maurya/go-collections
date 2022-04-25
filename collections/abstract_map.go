package collections

import (
	"errors"
)

var (
	InvalidKey = errors.New("invalid key")
)

type abstractMap[K comparable, V any] struct {
	data     map[K]V
	size     int
	nilValue V
}

func (al *abstractMap) Size() int {
	return al.size
}

func (al *abstractMap) IsEmpty() bool {
	return al.size == 0
}

func (al *abstractMap) Clear() {
	al.data = nil
	al.size = 0
}

func (al *abstractMap[K]) Contains(key K) bool {
	_, ok := al.data[key]
	return ok
}
