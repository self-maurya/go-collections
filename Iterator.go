package go_collections

type Iterator interface {
	HasNext() bool
	Next() interface{}
	Remove()

}
