package main

import (
	"fmt"
	"testing"
)

func (s String) Hash() uint {
	return djb2([]byte(s))
}

func TestDict(t *testing.T) {
	m := MakeDict[String, int](5)
	m.Insert("hello", 5)
	m.Insert("bye", 3)
	m.Insert("coosl", 55)
	m.Insert("dddd", 100)
	m.Insert("esdfg", 299)
	m.Insert("fsss", 444)
	m.Set("hello", 150)
	m.Insert("bye", 233)
	m.Pop("dddd")

	i := m.CreateIterator()
	for i.HasNext() {
		bucket := i.Next()
		fmt.Printf("%+v: %+v \n", bucket.key, bucket.value)
	}

	if false {
		t.Error("...")
	}
}
