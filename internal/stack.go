package internal

type Stack[T int | string] []T

func (s *Stack[T]) IsEmpty() bool {
	if s != nil && len(*s) > 0 {
		return false
	}
	return true
}

func (s *Stack[T]) Push(item T) {
	*s = append(*s, item)
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		panic("stack in empty")
	}

	last := len(*s) - 1
	element := (*s)[last]
	*s = (*s)[:last]

	return element
}

func (s *Stack[T]) Peek() T {
	if !s.IsEmpty() {
		last := len(*s) - 1
		element := (*s)[last]
		return element
	}

	var result T
	return result
}

func NewStack() *Stack[int] {
	return &Stack[int]{}
}

func NewStringStack() *Stack[string] {
	return &Stack[string]{}
}
