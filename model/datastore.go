package model

import (
	"fmt"
	"sync"

	"github.com/imhshekhar47/go-api-core/utils"
)

type Datastore[T Identifiable] struct {
	lock *sync.RWMutex

	list  *LinkedList[T]
	kvMap map[string]*Node[T]
}

func NewDatastore[T Identifiable](limit int32) *Datastore[T] {
	return &Datastore[T]{
		lock:  &sync.RWMutex{},
		list:  NewLinkedList[T](limit),
		kvMap: make(map[string]*Node[T]),
	}
}

func (ml *Datastore[T]) IsEmpty() bool {
	return ml.list.IsEmpty()
}

func (ml *Datastore[T]) Size() int32 {
	return ml.list.length
}

func (ml *Datastore[T]) Slice(from int32, upto int32) ([]T, error) {
	return ml.list.Slice(from, upto)
}

func (ml *Datastore[T]) FindById(key string) (*T, error) {
	ml.lock.Lock()
	defer ml.lock.Unlock()
	node, found := ml.kvMap[key]
	if !found {
		return nil, fmt.Errorf("item with key %s NOT FOUND", key)
	}

	return &node.data, nil
}

func (ml *Datastore[T]) Add(data T) error {
	if utils.IsEmpty(data.GetId()) {
		return fmt.Errorf("missing Key in the data")
	}
	node, found := ml.kvMap[data.GetId()]

	ml.lock.Lock()
	defer ml.lock.Unlock()
	if found {
		node.data = data
	} else {
		appendedNode, err := ml.list.Append(data)
		if err == nil {
			ml.kvMap[appendedNode.data.GetId()] = appendedNode
		} else {
			return err
		}
	}

	return nil
}

func (ml *Datastore[T]) Remove(key string) (*T, error) {
	node, found := ml.kvMap[key]

	if !found {
		return nil, fmt.Errorf("invalid key %s", key)
	}

	ml.lock.Lock()
	defer ml.lock.Unlock()
	ml.list.Remove(node.data)
	delete(ml.kvMap, node.data.GetId())

	return &node.data, nil
}
