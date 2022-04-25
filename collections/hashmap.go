package collections

type HashMap[K comparable, V any] struct {
	abstractMap[K, V]
}

func (hm *HashMap[K, V]) Add(key K, value V) {
	hm.data[key] = value
	hm.size++
}

func (hm *HashMap[K, V]) Get(key K) (V, error) {
	if hm.Contains(key) {
		return hm.data[key], nil
	}
	return nil, InvalidKey
}
