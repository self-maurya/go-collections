package collections

import (
	"errors"
)

var (
	InvalidIndexError = errors.New("invalid index")
)

type ArrayList[V comparable] struct {
	abstractList[V]
}

func (al *ArrayList[V]) checkIndex(index int) error {
	if index < 0 || index >= al.size {
		return InvalidIndexError
	}
	return nil
}

func (al *ArrayList[V]) Add(value V) {
	al.data = append(al.data, value)
	al.size++
}

func (al *ArrayList[V]) Get(index int) (V, error) {
	if err := al.checkIndex(index); err != nil {
		return al.nilValue, err
	}
	return al.data[index], nil
}

func (al *ArrayList[V]) Set(index int, value V) (V, error) {
	oldValue, err := al.Get(index)
	if err != nil {
		return al.nilValue, err
	}
	al.data[index] = value
	return oldValue, nil
}

func (al *ArrayList[V]) AddAtIndex(index int, value V) error {
	if err := al.checkIndex(index); err != nil {
		return err
	}
	prev := al.data[:index]
	next := al.data[index:]
	al.data = append(append(prev, value), next...)
	al.size++
	return nil
}

func (al *ArrayList[V]) Remove(index int) (V, error) {
	oldValue, err := al.Get(index)
	if err != nil {
		return al.nilValue, err
	}
	prev := al.data[:index]
	next := al.data[index+1:]
	al.data = append(prev, next...)
	al.size--
	return oldValue, nil
}

func (al *ArrayList[V]) IndexOf(value V) int {
	for i := 0; i < al.size; i++ {
		if al.data[i] == value {
			return i
		}
	}
	return -1
}

func (al *ArrayList[V]) Contains(value V) bool {
	return al.IndexOf(value) != -1
}

func (al *ArrayList[V]) LastIndexOf(value V) int {
	for i := al.size - 1; i >= 0; i-- {
		if al.data[i] == value {
			return i
		}
	}
	return -1
}

func (al *ArrayList[V]) SubList(fromIndex, toIndex int) ([]V, error) {
	if fromIndex < 0 || toIndex > al.size || fromIndex > toIndex {
		return nil, InvalidIndexError
	}
	return al.data[fromIndex:toIndex], nil
}

func (al *ArrayList[V]) RemoveRange(fromIndex, toIndex int) error {
	if fromIndex < 0 || toIndex > al.size || fromIndex > toIndex {
		return InvalidIndexError
	}
	prev := al.data[:fromIndex]
	next := al.data[toIndex:]
	al.data = append(prev, next...)
	al.size -= toIndex - fromIndex
	return nil
}
