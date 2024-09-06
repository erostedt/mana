package mana

import (
	"testing"
)

func (s String) Hash() uint {
	return djb2([]byte(s))
}

func TestDictInsertResize(t *testing.T) {
	d := MakeDict[String, int]()
	d.Insert("a", 1)
	d.Insert("b", 2)
	d.Insert("c", 3)
	d.Insert("d", 4)
	d.Insert("e", 5)
	d.Insert("f", 6)
	d.Insert("g", 7)
	d.Insert("h", 8)
	d.Insert("i", 9)

	if d.size != 9 {
		t.Error("Dict should have size of 6.")
	}
	if d.cap != 18 {
		t.Error("Dict should have cap of 18.")
	}
}

func TestDictGet(t *testing.T) {
	d := MakeDict[String, int]()
	d.Insert("a", 1)

	_, err := d.Get("a")
	if err != nil {
		t.Error("Could not get a")
	}
}

func TestDictPop(t *testing.T) {
	d := MakeDict[String, int]()
	d.Insert("a", 1)
	d.Insert("b", 2)

	{
		value, err := d.Pop("b")
		if err != nil {
			t.Error("Could not pop b")
		}
		t.Logf("Popped %d", value)
	}
	{
		_, err := d.Pop("b")
		if err == nil {
			t.Error("Could pop b multiple times")
		}
	}
	{
		if d.Contains("b") {
			t.Error("Should not contain b")
		}
		if !d.Contains("a") {
			t.Error("Should contain a")
		}
	}
}
