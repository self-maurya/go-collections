package src

import (
	"fmt"
	"testing"
)

func TestArrayList_Add(t *testing.T) {
	al := &ArrayList[string]{}
	al.Add("a")
	al.Add("c")
	al.Add("d")
	_ = al.AddAtIndex(1, "b")
	for _, v := range al.ToArray() {
		fmt.Println(v)
	}
}
