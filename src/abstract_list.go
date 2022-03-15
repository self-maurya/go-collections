package src

type abstractList[V comparable] struct {
	elementData []V
	size        int
	nilValue    V
}

func (al *abstractList[V]) Size() int {
	return al.size
}

func (al *abstractList[V]) IsEmpty() bool {
	return al.size == 0
}

func (al *abstractList[V]) Clear() {
	al.elementData = nil
	al.size = 0
}

func (al *abstractList[V]) ToArray() []V {
	return al.elementData
}

func (al *abstractList[V]) AddAll(value []V) {
	al.elementData = append(al.elementData, value...)
	al.size += len(value)
}

func (al *abstractList[V]) Equals(list *abstractList[V]) bool {
	if list.size != al.size {
		return false
	}
	for i := 0; i < al.size; i++ {
		if list.elementData[i] != al.elementData[i] {
			return false
		}
	}
	return true
}
