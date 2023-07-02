package main

import (
	"testing"
)

func (s String) Hash() uint {
	return djb2([]byte(s))
}

func TestDictInsertResize(t *testing.T) {
	d := MakeDict[String, int](5)
	d.Insert("a", 1)
	d.Insert("b", 2)
	d.Insert("c", 3)
	d.Insert("d", 4)
	d.Insert("e", 5)
	d.Insert("f", 6)

	if d.size != 6 {
		t.Error("Dict should have size of 6.")
	}

	if d.cap != 12 {
		t.Errorf("Dict should have size of 12, had %d", d.cap)
	}
}
