package linkedlist

import (
	"errors"
	"fmt"
)

type Node[T comparable] struct {
	value T
	next *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
	length uint
}

func (l *LinkedList[T]) Init() {
	l.head = nil
	l.length = 0
}

func (l *LinkedList[T]) Destroy() error {
	if l.head != nil || l.length != 0 {
		return errors.New("cannot destroy non-empty list")
	}

	return nil
}

func (l *LinkedList[T]) Prepend(value T) {
	node := Node[T] {
		value: value,
		next: nil,
	}

	l.length += 1
	node.next = l.head
	l.head = &node
}

func (l *LinkedList[T]) Append(value T) {
	if l.head == nil {
		l.Prepend(value)
		return
	}

	node := Node[T] {
		value: value,
		next: nil,
	}

	clone := l.head
	for ; clone.next != nil; {
		clone = clone.next
	}

	l.length += 1
	clone.next = &node
}

func (l *LinkedList[T]) InsertAt(idx uint, value T) error {
	if idx > l.length {
		return fmt.Errorf(fmt.Sprintf("index: %v out of bounds", idx))
	} else if idx == 0 {
		l.Prepend(value)
	} else if idx == l.length {
		l.Append(value)
	}

	node := Node[T] {
		value: value,
		next: nil,
	}

	clone := l.head
	for i := 1; i < int(idx); i++ {
		clone = clone.next
	}

	node.next = clone.next
	clone.next = &node
	l.length += 1
	
	return nil
}

func (l *LinkedList[T]) Get(idx uint) (value T, err error) {
	if idx >= l.length {
		err = fmt.Errorf(fmt.Sprintf("index: %v out of bounds", idx))
		return
	} else if idx == 0 {
		value = l.head.value
		err =  nil
		return
	}

	clone := l.head
	for i := 1; i <= int(idx); i++ {
		clone = clone.next
	}

	value = clone.value
	err = nil
	return
}

func (l *LinkedList[T]) RemoveAt(idx uint) (value T, err error) {
	if idx >= l.length {
		err = fmt.Errorf(fmt.Sprintf("index: %v out of bounds", idx))
		return
	} else if idx == 0 {
		node := l.head
		l.head = l.head.next
		l.length -= 1

		node.next = nil
		value = node.value
		err = nil
		return
	}

	clone := l.head
	for i := 0; i < int(idx); i++ {
		clone = clone.next
	}

	node := clone.next
	clone.next = clone.next.next
	l.length -= 1

	node.next = nil
	value = node.value
	err = nil

	return
}

func (l *LinkedList[T]) Remove(value T) error {
	if l.head == nil {
		return errors.New("cannot remove from an empty list")
	} else if value == l.head.value {
		_, err := l.RemoveAt(0)
		return err
	}

	clone := l.head
	for i := 1; i < int(l.length); i++ {
		clone = clone.next
		if value == clone.value {
			_, err := l.RemoveAt(uint(i))
			return err
		}
	}

	return errors.New("value not found")
}