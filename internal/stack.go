package internal

type Stack []int

func (s *Stack) IsEmpty() bool {
	if s != nil && len(*s) > 0 {
		return false
	}
	return true
}

func (s *Stack) Push(d int) {
	*s = append(*s, d)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("stack in empty")
	}

	last := len(*s) - 1
	element := (*s)[last]
	*s = (*s)[:last]

	return element
}

func (s *Stack) Peek() int {
	if !s.IsEmpty() {
		last := len(*s) - 1
		element := (*s)[last]
		return element
	}

	return 0
}

func NewStack() *Stack {
	return &Stack{}
}
