package collections

import (
	"errors"
)

var (
	StackIsEmpty = errors.New("stack is empty")
)

type Stack[V comparable] struct {
	abstractList[V]
}

func (s *Stack[V]) Push(value V) {
	s.data = append(s.data, value)
	s.size++
}

func (s *Stack[V]) Pop() (V, error) {
	if s.IsEmpty() {
		return s.nilValue, StackIsEmpty
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	s.size--
	return top, nil
}
