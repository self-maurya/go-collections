package collections

import "testing"

func TestTreeMap_Add(t *testing.T) {
	tm := TreeMap[int, string]{}
	tm.Add(1, "a")
	tm.Add(2, "b")
	tm.Add(3, "c")
	if tm.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", tm.Size())
	}
}

func TestTreeMap_Get(t *testing.T) {
	tm := TreeMap[int, string]{}
	tm.Add(1, "a")
	v, err := tm.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if v != "a" {
		t.Errorf("Expected value to be a, got %s", v)
	}
}

func TestTreeMap_Clear(t *testing.T) {
	tm := TreeMap[int, string]{}
	tm.Add(1, "a")
	tm.Clear()
	if !tm.IsEmpty() {
		t.Errorf("Expected IsEmpty to be true, got false")
	}
}

func TestTreeMap_Contains(t *testing.T) {
	tm := TreeMap[int, string]{}
	tm.Add(1, "a")
	if !tm.Contains(1) {
		t.Errorf("Expected Contains to be true, got false")
	}
}
