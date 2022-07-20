package model

import "fmt"

type LinkedList[T Identifiable] struct {
	limit  int32
	length int32

	start *Node[T]
	end   *Node[T]
}

func NewLinkedList[T Identifiable](limit int32) *LinkedList[T] {
	return &LinkedList[T]{
		limit: limit,
	}
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.start == nil
}

func (l *LinkedList[T]) Append(data T) (*Node[T], error) {
	if l.limit == l.length {
		return nil, fmt.Errorf("list overflow")
	}

	node := NewNode(data)
	if l.start == nil {
		l.start = node
	} else {
		l.end.next = node
		node.prev = l.end
	}
	l.end = node

	l.length = l.length + 1
	return node, nil
}

func (l *LinkedList[T]) Remove(data T) (*Node[T], error) {
	if l.start == nil {
		return nil, fmt.Errorf("list is empty")
	}

	isFound := false
	runner := l.start
	for runner != nil {
		if runner.data.GetId() == data.GetId() {
			isFound = true
			break
		}
		runner = runner.next
	}

	if !isFound {
		return nil, fmt.Errorf("data with key %s not found", data.GetId())
	}

	next := runner.next
	prev := runner.prev

	if next == nil && prev == nil {
		l.start = nil
		l.end = nil
	} else if next == nil {
		prev.next = next
		l.end = prev
	} else if prev == nil {
		next.prev = prev
		l.start = next
	} else {
		prev.next = next
		next.prev = prev
	}
	l.length = l.length - 1
	return runner, nil
}

func (l *LinkedList[T]) Slice(from int32, upto int32) ([]T, error) {
	if l.IsEmpty() || l.length < from {
		return nil, fmt.Errorf("out of range %d to %d", from, upto)
	}

	begin := l.start
	var index int32 = 0
	for ; index < from; index++ {
		if begin.next != nil {
			begin = begin.next
		}
	}

	result := make([]T, 0)
	for ; index < upto; index++ {
		result = append(result, begin.data)
		if begin.next == nil {
			break
		}

		begin = begin.next
	}

	return result, nil
}
