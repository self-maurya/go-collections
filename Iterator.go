package go_collections

type Iterator interface {
	HasNext() bool
	Next[V comparable]() V
	Remove()
}
