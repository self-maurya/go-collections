package src

import (
	"errors"
)

var (
	StackIsEmpty = errors.New("stack is empty")
)

type Stack[V comparable] struct {
	abstractList[V]
}

func (s *Stack[V]) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack[V]) Push(value V) {
	s.elementData = append(s.elementData, value)
	s.size++
}

func (s *Stack[V]) Pop() (V, error) {
	if s.IsEmpty() {
		return s.nilValue, StackIsEmpty
	}
	top := s.elementData[len(s.elementData)-1]
	s.elementData = s.elementData[:len(s.elementData)-1]
	s.size--
	return top, nil
}
