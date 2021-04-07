package go_collections

type List interface {
	Collection
	Get(int) (interface{}, error)
	Set(int, interface{}) (interface{}, error)
	IndexOf(interface{}) int
	LastIndexOf(interface{}) int
	SubList(int, int) List
}
