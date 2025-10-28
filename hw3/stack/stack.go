package stack

type Stack[T any] struct {
	data []T
}

func New[T any](capHint ...int) *Stack[T] {
	if len(capHint) > 0 && capHint[0] > 0 {
		return &Stack[T]{data: make([]T, 0, capHint[0])}
	}
	return &Stack[T]{}
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	i := len(s.data) - 1
	v := s.data[i]
	s.data[i] = zero
	s.data = s.data[:i]
	return v, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) Clear() {
	var zero T
	for i := range s.data {
		s.data[i] = zero
	}
	s.data = s.data[:0]
}
