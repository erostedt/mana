package mana

import (
	"errors"
	"fmt"
)

type DequeNode[T any] struct {
	data T
	prev *DequeNode[T]
	next *DequeNode[T]
}

type Deque[T any] struct {
	head  *DequeNode[T]
	tail  *DequeNode[T]
	count uint
}

func MakeDeque[T any]() Deque[T] {
	return Deque[T]{nil, nil, 0}
}

func NewDequeNode[T any](data T) *DequeNode[T] {
	node := new(DequeNode[T])
	node.data = data
	node.prev = nil
	node.next = nil
	return node
}

func (q *Deque[T]) AddFirst(data T) {
	node := NewDequeNode(data)
	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		node.next = q.head
		q.head = node
		q.head.prev = q.head
	}
}

func (q *Deque[T]) AddLast(data T) {
	node := NewDequeNode(data)
	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		node.prev = q.tail
		q.tail = node
	}
	q.count++
}

func (q *Deque[T]) PopFirst() (data T, err error) {
	if q.count == 0 {
		return data, errors.New("Deque is empty")
	}

	if q.count == 1 {
		data = q.head.data
		q.head = nil
		q.tail = nil
		q.count--
		return data, nil
	}

	q.count--
	data = q.head.data
	q.head = q.head.next
	q.head.prev = nil
	return data, nil
}

func (q *Deque[T]) PopLast() (data T, err error) {
	if q.count == 0 {
		return data, errors.New("Deque is empty")
	}

	if q.count == 1 {
		data = q.head.data
		q.head = nil
		q.tail = nil
		q.count--
		return data, nil
	}

	q.count--
	data = q.tail.data
	q.tail.prev.next = nil
	q.tail = q.tail.prev
	return data, nil
}

func (q *Deque[T]) IsEmpty() bool {
	return q.count == 0
}

func (q *Deque[T]) PrintDeque() {
	node := q.head
	if node == nil {
		return
	}

	for node.next != nil {
		fmt.Print(node.data)
		fmt.Print(" <-> ")
		node = node.next
	}
	fmt.Println(node.data)
}
