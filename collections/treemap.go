package collections

type TreeMap[K comparable, V any] struct {
	data map[K]V
	size int
	nilValue V
}

func (tm *TreeMap[K, V]) Add(key K, value V) {
	tm.data[key] = value
	tm.size++
}

func (tm *TreeMap[K, V]) Get(key K) (V, error) {
	if tm.Contains(key) {
		return tm.data[key], nil
	}
	return tm.nilValue, InvalidKey
}

func (tm *TreeMap[K, V]) Size() int{
	return tm.size
}

func (tm *TreeMap[K, V]) IsEmpty() bool{
	return tm.size == 0
}

func (tm *TreeMap[K, V]) Clear(){
	tm.data = make(map[K]V)
	tm.size = 0
}
func (tm *TreeMap[K, V]) Contains(key K) bool{
	_,ok := tm.data[key]
	return ok
}
