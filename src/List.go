package src

type List[V comparable] interface {
	Collection[V]
	Get(int) (V, error)
	Set(int, V) (V, error)
	IndexOf(V) int
	LastIndexOf(V) int
	SubList(int, int) ([]V, error)
}
