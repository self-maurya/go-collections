package collections

type abstractList[V comparable] struct {
	data     []V
	size     int
	nilValue V
}

func (al *abstractList[V]) Size() int {
	return al.size
}

func (al *abstractList[V]) IsEmpty() bool {
	return al.size == 0
}

func (al *abstractList[V]) Clear() {
	al.data = nil
	al.size = 0
}

func (al *abstractList[V]) ToArray() []V {
	return al.data
}

func (al *abstractList[V]) AddAll(value []V) {
	al.data = append(al.data, value...)
	al.size += len(value)
}

func (al *abstractList[V]) Equals(list *abstractList[V]) bool {
	if list.size != al.size {
		return false
	}
	for i := 0; i < al.size; i++ {
		if list.data[i] != al.data[i] {
			return false
		}
	}
	return true
}
