package collections

type HashSet[V comparable] struct {
	abstractMap[V, interface{}]
}

func (hm *HashSet[V]) Add(value V) {
	hm.data[value] = struct{}{}
	hm.size++
}

func (hm *HashSet[V]) Remove(value V) {
	delete(hm.data, value)
}
