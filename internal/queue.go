package internal

type Queue[T int | string] []T

func (q *Queue[T]) IsEmpty() bool {
	if q != nil && len(*q) > 0 {
		return false
	}
	return true
}

func (q *Queue[T]) Enqueue(i T) {
	*q = append(*q, i)
}

func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element
}

func NewQueue() *Queue[int] {
	return &Queue[int]{}
}

func NewStringQueue() *Queue[string] {
	return &Queue[string]{}
}
