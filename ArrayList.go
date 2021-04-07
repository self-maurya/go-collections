package go_collections

import (
	"errors"
)

type ArrayList struct {
	elementData []interface{}
	size int
}

func (al *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= al.size {
		return nil, errors.New("invalid index")
	}
	return al.elementData[index], nil
}

func (al *ArrayList) Set(index int, value interface{}) (interface{}, error) {
	if index < 0 || index >= al.size {
		return nil, errors.New("invalid index")
	}
	oldValue := al.elementData[index]
	al.elementData[index] = value
	return oldValue, nil
}



