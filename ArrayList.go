package go_collections

import (
	"errors"
)

var (
	InvalidIndexError = errors.New("invalid index")
)

type ArrayList[V comparable] struct {
	elementData []V
	size        int
	nilValue    V
}

func (al *ArrayList[V]) Add(value V) {
	al.elementData = append([]V(al.elementData), value)
	al.size++
}

func (al *ArrayList[V]) Get(index int) (V, error) {
	if index < 0 || index >= al.size {
		return al.nilValue, InvalidIndexError
	}
	return al.elementData[index], nil
}

func (al *ArrayList[V]) Set(index int, value V) (V, error) {
	if index < 0 || index >= al.size {
		return al.nilValue, InvalidIndexError
	}
	oldValue := al.elementData[index]
	al.elementData[index] = value
	return oldValue, nil
}

func (al *ArrayList[V]) AddAtIndex(index int, value V) error {
	if index < 0 || index >= al.size {
		return InvalidIndexError
	}
	prev := al.elementData[:index]
	next := al.elementData[index:]
	al.elementData = append(append([]V(prev), value), next...)
	al.size++
	return nil
}

func (al *ArrayList[V]) Remove(index int) (V, error) {
	if index < 0 || index >= al.size {
		return al.nilValue, InvalidIndexError
	}
	oldValue := al.elementData[index]
	prev := al.elementData[:index]
	next := al.elementData[index+1:]
	al.elementData = append([]V(prev), next...)
	al.size--
	return oldValue, nil
}

func (al *ArrayList[V]) IndexOf(value V) int {
	for i := 0; i < al.size; i++ {
		if al.elementData[i] == value {
			return i
		}
	}
	return -1
}

func (al *ArrayList[V]) LastIndexOf(value V) int {
	for i := al.size - 1; i >= 0; i-- {
		if al.elementData[i] == value {
			return i
		}
	}
	return -1
}

func (al *ArrayList[V]) Clear(value V) int {
	al.elementData = nil
	al.size = 0
	return -1
}

func (al *ArrayList[V]) AddAll(value []V) {
	al.elementData = append([]V(al.elementData), value...)
	al.size += len(value)
}

func (al *ArrayList[V]) SubList(fromIndex, toIndex int) ([]V, error) {
	if fromIndex < 0 || toIndex > al.size || fromIndex > toIndex {
		return nil, InvalidIndexError
	}
	return al.elementData[fromIndex:toIndex], nil
}

func (al *ArrayList[V]) Equals(list *ArrayList[V]) bool {
	if list.size != al.size {
		return false
	}
	for i := 0; i < al.size; i++ {
		if list.elementData[i] != al.elementData[i] {
			return false
		}
	}
	return true
}

func (al *ArrayList[V]) RemoveRange(fromIndex, toIndex int) error {
	if fromIndex < 0 || toIndex > al.size || fromIndex > toIndex {
		return InvalidIndexError
	}
	prev := al.elementData[:fromIndex]
	next := al.elementData[toIndex:]
	al.elementData = append([]V(prev), next...)
	al.size -= toIndex - fromIndex
	return nil
}

func (al *ArrayList[V]) Iterator() []V {
	return al.elementData
}
