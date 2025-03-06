package collections

import "errors"

type LinkedList[V comparable] struct {
	head *node[V]
	tail *node[V]
	size int
	nilValue V
}

type node[V comparable] struct {
	value V
	next  *node[V]
	prev *node[V]
}

func NewLinkedList[V comparable]() *LinkedList[V]{
	return &LinkedList[V]{nil,nil,0,nilValue}
}

func (l *LinkedList[V]) Add(value V){
	n := &node[V]{value,nil,nil}
	if l.head == nil{
		l.head = n
		l.tail = n
	}else{
		l.tail.next = n
		n.prev = l.tail
		l.tail = n
	}
	l.size++
}

func (l *LinkedList[V]) AddAll(values []V){
	for _,v := range values{
		l.Add(v)
	}
}

func (l *LinkedList[V]) Get(index int) (V,error){
	if index < 0 || index >= l.size{
		return l.nilValue, errors.New("Invalid index")
	}
	curr := l.head
	for i := 0; i < index; i++ {
		if l.head.value == value{
			return l.head.value,nil
		}
		if l.head.next != nil{
			curr = l.head.next
		}
		if l.head.prev != nil{
			curr = l.head.prev
		}
		
	}
	return l.nilValue, errors.New("Invalid index")
}

func (l *LinkedList[V]) Remove(index int) (V,error){
	if index < 0 || index >= l.size{
		return l.nilValue, errors.New("Invalid index")
	}
	curr := l.head
	for i := 0; i < index; i++ {
		if l.head.value == value{
			l.head = l.head.next
			return l.head.value,nil
		}
		if l.head.next != nil{
			l.head.next.prev = nil
			l.head = l.head.next
		}
		if l.head.prev != nil{
			l.head.prev.next = nil
			l.head = l.head.prev
		}
		
	}
	return l.nilValue, errors.New("Invalid index")
}

func (l *LinkedList[V]) Size() int{
	return l.size
}

func (l *LinkedList[V]) IsEmpty() bool{
	return l.size == 0
}

func (l *LinkedList[V]) Clear(){
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *LinkedList[V]) Contains(value V) bool{
	curr := l.head
	for i:=0;i<l.size;i++{
		if curr.value == value{
			return true
		}
		if curr.next != nil{
			curr = curr.next
		}
		if curr.prev != nil{
			curr = curr.prev
		}
	}
	return false
}

func (l *LinkedList[V]) ToArray() []V{
	var values []V
	curr := l.head
	for i:=0;i<l.size;i++{
		values = append(values,curr.value)
		curr = curr.next
	}
	return values
}

func (l *LinkedList[V]) Equals(other *LinkedList[V]) bool{
	if l.size != other.size{
		return false
	}
	curr := l.head
	for i:=0;i<l.size;i++{
		if curr.value != other.head.value{
			return false
		}
		if curr.next != nil{
			curr = curr.next
		}
		if curr.prev != nil{
			curr = curr.prev
		}
	}
	return true
}

func (l *LinkedList[V]) IndexOf(value V) int {
	for i := 0; i < l.size; i++ {
		if l.head.value == value {
			return i
		}
		if l.head.next != nil{
			l.head = l.head.next
		}
		if l.head.prev != nil{
			l.head = l.head.prev
		}
	}
	return -1
}

func (l *LinkedList[V]) LastIndexOf(value V) int{
	for i := l.size - 1; i >= 0; i-- {
		if l.head.value == value {
			return i
		}
		if l.head.next != nil{
			l.head = l.head.next
		}
		if l.head.prev != nil{
			l.head = l.head.prev
		}
	}
	return -1
}
