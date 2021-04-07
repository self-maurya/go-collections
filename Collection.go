package go_collections

type Collection interface {
	Iterable
	Size() int
	IsEmpty() bool
	Contains(interface{}) bool
	ToArray() []interface{}
	Add(...interface{}) bool
	Remove(interface{}) bool
	Clear()
}
