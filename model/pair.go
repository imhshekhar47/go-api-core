package model

type Pair[T1 any, T2 any] struct {
	first  T1
	second T2
}

func NewPair[T1 any, T2 any](first T1, second T2) *Pair[T1, T2] {
	return &Pair[T1, T2]{
		first:  first,
		second: second,
	}
}

func (s *Pair[T1, T2]) First() T1 {
	return s.first
}

func (s *Pair[T1, T2]) Second() T2 {
	return s.second
}
