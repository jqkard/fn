package ds

type Stack[T any] struct {
	items []T
}

func NewStack[T any](capacity int) *Stack[T] {
	return &Stack[T]{items: make([]T, capacity)}
}

func StackFrom[T any](items []T) *Stack[T] {
	s := NewStack[T](0)
	s.items = items
	return s
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	topIndex := len(s.items) - 1
	top := s.items[topIndex]
	s.items = s.items[:topIndex]
	return top
}

func (s Stack[T]) Top() T {
	topIndex := len(s.items) - 1
	return s.items[topIndex]
}

func (s Stack[T]) Len() int {
	return len(s.items)
}

func (s Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s Stack[T]) Items() []T {
	return s.items
}
