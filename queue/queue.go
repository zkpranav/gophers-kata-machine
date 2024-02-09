package queue

import "errors"

type Node[T any] struct {
	value T
	next *Node[T]
}

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func (q *Queue[T]) Init() {
	q.head = nil
	q.tail = nil
}

func (q *Queue[T]) Destroy() error {
	if q.head != nil || q.tail != nil {
		return errors.New("non-empty queue cannot be destroyed")
	}

	return nil
}

func (q *Queue[T]) Enqueue(value T) {
	node := Node[T] {
		value: value,
		next: nil,
	}

	if (q.tail == nil) {
		q.tail = &node
		return
	}

	q.tail.next = &node
	q.tail = &node
}

func (q *Queue[T]) Deque() (value T, err error) {
	if q.head == nil {
		err = errors.New("cannot deque an empty queue")
		return
	}

	node := q.head
	q.head = q.head.next
	node.next = nil
	value = node.value
	return
}

func (q *Queue[T]) Peek() (value T) {
	if q.head == nil {
		return
	}

	return q.head.value
}