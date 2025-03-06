package collections

type TreeSet[V comparable] struct {
	data map[V]interface{}
	size int
	nilValue V
}

func (ts *TreeSet[V]) Add(value V) {
	ts.data[value] = struct{}{}
	ts.size++
}

func (ts *TreeSet[V]) Remove(value V){
	delete(ts.data, value)
}

func (ts *TreeSet[V]) Size() int{
	return ts.size
}

func (ts *TreeSet[V]) IsEmpty() bool{
	return ts.size == 0
}

func (ts *TreeSet[V]) Clear(){
	ts.data = make(map[V]interface{})
	ts.size = 0
}
func (ts *TreeSet[V]) Contains(value V) bool{
	_,ok := ts.data[value]
	return ok
}
