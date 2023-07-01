package main

import (
	"testing"
)

func TestDeque(t *testing.T) {
	q := MakeDeque[int]()
	q.AddLast(1)
	q.AddLast(2)
	q.AddLast(3)

	q.PrintDeque()

	q.PopFirst()
	q.PopFirst()

	q.PrintDeque()

	if false {
		t.Error("...")
	}
}
