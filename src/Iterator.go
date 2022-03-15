package src

type Iterator[V comparable] interface {
	HasNext() bool
	Next() V
	Remove()
}
