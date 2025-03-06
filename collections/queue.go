package collections

import "errors"

type Queue[V comparable] struct {
	data []V
	size int
	nilValue V
}

func (q *Queue[V]) Enqueue(value V) {
	q.data = append(q.data, value)
	q.size++
}

func (q *Queue[V]) Dequeue() (V, error) {
	if q.IsEmpty() {
		return q.nilValue, errors.New("queue is empty")
	}
	front := q.data[0]
	q.data = q.data[1:]
	q.size--
	return front, nil
}

func (q *Queue[V]) Size() int {
	return q.size
}

func (q *Queue[V]) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue[V]) Clear() {
	q.data = nil
	q.size = 0
}

func (q *Queue[V]) Contains(value V) bool {
	for _, v := range q.data {
		if v == value {
			return true
		}
	}
	return false
}

func (q *Queue[V]) ToArray() []V {
	return q.data
}

func (q *Queue[V]) AddAll(values []V) {
	for _, v := range values {
		q.Enqueue(v)
	}
}
