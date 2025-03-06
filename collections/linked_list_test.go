package collections

import "testing"

func TestLinkedList_Add(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	if ll.Size() != 3{
		t.Errorf("Expected size to be 3, got %d",ll.Size())
	}
}

func TestLinkedList_AddAll(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.AddAll([]int{1,2,3})
	if ll.Size() != 3{
		t.Errorf("Expected size to be 3, got %d",ll.Size())
	}
}

func TestLinkedList_Get(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Add(1)
	v,err := ll.Get(0)
	if err != nil{
		t.Errorf("Expected no error, got %s",err)
	}
	if v != 1{
		t.Errorf("Expected value to be 1, got %d",v)
	}
}

func TestLinkedList_Remove(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Add(1)
	v,err := ll.Remove(0)
	if err != nil{
		t.Errorf("Expected no error, got %s",err)
	}
	if v != 1{
		t.Errorf("Expected value to be 1, got %d",v)
	}
}

func TestLinkedList_Clear(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Add(1)
	ll.Clear()
	if !ll.IsEmpty(){
		t.Errorf("Expected IsEmpty to be true, got false")
	}
}

func TestLinkedList_Contains(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Add(1)
	if !ll.Contains(1){
		t.Errorf("Expected Contains to be true, got false")
	}
}
