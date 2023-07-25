package main

import (
	"testing"
)

func TestDequeAddAndRemove(t *testing.T) {
	q := MakeDeque[int]()
	q.AddLast(1)
	q.AddLast(2)
	q.AddLast(3)

	a, _ := q.PopFirst()
	if a != 1 {
		t.Error("First popped element exptected to be 1.")
	}

	b, _ := q.PopLast()
	if b != 3 {
		t.Error("PopLast expected to return 3.")
	}

	if q.count != 1 {
		t.Error("Count expected to be 1.")
	}

	if q.head != q.tail {
		t.Error("Head expected to be equal to tail.")
	}

	q.PopFirst()

	if q.head != nil {
		t.Error("Head should be nil.")
	}

	if q.count != 0 || !q.IsEmpty() {
		t.Error("Count should be 0.")
	}

	if q.tail != nil {
		t.Error("Tail should be nil.")
	}

	q.AddFirst(4)
	q.AddFirst(5)
	if q.head.data != 5 {
		t.Error("Head should be 5")
	}

	if q.head.next.data != 4 || q.tail.data != 4 {
		t.Error("Tail should be 4")
	}
}
