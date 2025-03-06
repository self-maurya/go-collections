package collections

import "testing"

func TestQueue_Enqueue(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if q.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", q.Size())
	}
}

func TestQueue_Dequeue(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(1)
	v, err := q.Dequeue()
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if v != 1 {
		t.Errorf("Expected value to be 1, got %d", v)
	}
}

func TestQueue_Clear(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(1)
	q.Clear()
	if !q.IsEmpty() {
		t.Errorf("Expected IsEmpty to be true, got false")
	}
}

func TestQueue_Contains(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(1)
	if !q.Contains(1) {
		t.Errorf("Expected Contains to be true, got false")
	}
}
