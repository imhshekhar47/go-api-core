package model

type Node[T Identifiable] struct {
	next *Node[T]
	prev *Node[T]
	data T
}

func NewNode[T Identifiable](data T) *Node[T] {
	return &Node[T]{
		data: data,
		next: nil,
		prev: nil,
	}
}
