package go_collections

type Iterator[V comparable] interface {
	HasNext() bool
	Next() V
	Remove()
}
