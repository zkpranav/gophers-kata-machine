/*
* Stack - 
* A LIFO data structure implemented with a singly linked list.
*
* Worst case time complexity - 
* Push: O(1)
* Pop: O(1)
* Peek: O(1)
*/

package stack

import "errors"

type Node[T any] struct {
	value T
	next *Node[T]
}

type Stack[T any] struct {
	head *Node[T]
}

func (s *Stack[T]) Init() {
	s.head = nil
}

func (s *Stack[T]) Destroy() error {
	if s.head != nil {
		return errors.New("non-empty stack cannot be destroyed")
	}

	return nil
}

func (s *Stack[T]) Push(value T) {
	node := Node[T] {
		value: value,
		next: nil,
	}

	node.next = s.head
	s.head = &node
}

func (s *Stack[T]) Pop() (value T, err error) {
	if s.head == nil {
		err = errors.New("cannot pop an empty stack")
		return
	}

	node := s.head
	s.head = s.head.next
	node.next = nil
	value = node.value
	return
}

func (s *Stack[T]) Peek() (value T) {
	if s.head == nil {
		return
	}

	return s.head.value
}
