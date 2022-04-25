package collections

type Collection[V comparable] interface {
	Size() int
	IsEmpty() bool
	Contains(V) bool
	ToArray() []V
	AddAll([]V)
	Remove(int) (V, error)
	Clear()
}
