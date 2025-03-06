package collections

import "testing"

func TestTreeSet_Add(t *testing.T) {
	ts := TreeSet[int]{}
	ts.Add(1)
	ts.Add(2)
	ts.Add(3)
	if ts.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", ts.Size())
	}
}

func TestTreeSet_Remove(t *testing.T) {
	ts := TreeSet[int]{}
	ts.Add(1)
	ts.Remove(1)
	if ts.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", ts.Size())
	}
}

func TestTreeSet_Clear(t *testing.T) {
	ts := TreeSet[int]{}
	ts.Add(1)
	ts.Clear()
	if !ts.IsEmpty() {
		t.Errorf("Expected IsEmpty to be true, got false")
	}
}

func TestTreeSet_Contains(t *testing.T) {
	ts := TreeSet[int]{}
	ts.Add(1)
	if !ts.Contains(1) {
		t.Errorf("Expected Contains to be true, got false")
	}
}
