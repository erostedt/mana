package main

import (
	"errors"
	"fmt"
)

type String string

type Hashable interface {
	Hash() uint
}

type Key interface {
	Hashable
	comparable
}

type HashMap[key Key, value any] struct {
	keys        []key
	values      []value
	occupied    []bool
	size        uint
	cap         uint
	extendLimit uint
}

func (m *HashMap[Key, Value]) Init(cap uint) *HashMap[Key, Value] {
	m.size = 0
	m.extendLimit = cap / 2
	m.cap = cap

	m.keys = make([]Key, cap)
	m.values = make([]Value, cap)
	m.occupied = make([]bool, cap)
	return m
}

func (m *HashMap[Key, Value]) Insert(key Key, value Value) bool {
	if m.size >= m.extendLimit {
		m.Extend()
	}
	return m._Insert(key, value)
}

func (m *HashMap[Key, Value]) Extend() {
	keys := m.keys
	values := m.values
	occupied := m.occupied

	m.size = 0
	m.extendLimit = m.cap + 1
	m.cap = (m.cap + 1) * 2
	m.keys = make([]Key, m.cap)
	m.values = make([]Value, m.cap)
	m.occupied = make([]bool, m.cap)

	for keyIndex, key := range keys {
		if occupied[keyIndex] {
			m._Insert(key, values[keyIndex])
		}
	}
}

func (m *HashMap[Key, Value]) _Insert(key Key, value Value) bool {
	slot := key.Hash() % m.cap
	var i uint = 0
	for ; i < m.cap; i++ {
		if m.keys[slot] == key {
			return false
		}
		if !m.occupied[slot] {
			m.keys[slot] = key
			m.values[slot] = value
			m.occupied[slot] = true
			m.size++
			return true
		}
		slot = (slot + 1) % m.cap
	}
	return false
}

func (m *HashMap[Key, Value]) Get(key Key) (Value, error) {
	slot, err := m._FindSlot(key)
	if err == nil {
		return m.values[slot], nil
	}
	return *new(Value), errors.New("key not found")
}

func (m *HashMap[Key, Value]) GetDefault(key Key, def Value) Value {
	value, err := m.Get(key)
	if err == nil {
		return value
	}
	return def
}

func (m *HashMap[Key, Value]) Set(key Key, value Value) error {
	slot, err := m._FindSlot(key)
	if err == nil {
		m.values[slot] = value
		return nil
	}
	return err
}

func (m *HashMap[Key, Value]) SetInsert(key Key, value Value) {
	err := m.Set(key, value)
	if err == nil {
		m.Insert(key, value)
	}
}

func (m *HashMap[Key, Value]) Pop(key Key) (Value, error) {
	slot, err := m._FindSlot(key)
	if err == nil {
		value := m.values[slot]
		m.occupied[slot] = false
		return value, nil
	}
	return *new(Value), errors.New("could not pop item, since it was not present")
}

func (m *HashMap[Key, Value]) Contains(key Key) bool {
	_, err := m._FindSlot(key)
	return err == nil
}

func (m *HashMap[Key, Value]) _FindSlot(key Key) (uint, error) {
	slot := key.Hash() % m.cap
	var i uint = 0
	for ; i < m.cap; i++ {
		if !m.occupied[slot] {
			return 0, errors.New("key not found")
		}

		if m.keys[slot] == key {
			return slot, nil
		}

		slot = (slot + 1) % m.cap
	}
	return 0, errors.New("key not found")
}

func (m *HashMap[Key, Value]) Print() {
	for keyIndex, key := range m.keys {
		if m.occupied[keyIndex] {
			fmt.Printf("%+v: %+v \n", key, m.values[keyIndex])
		}
	}
}

func (s String) Hash() uint {
	return djb2([]byte(s))
}

func djb2(bytes []byte) uint {
	var hash uint = 5381

	for _, c := range bytes {
		hash = (((hash << 5) + hash) + uint(c))
	}
	return hash
}

func a() {
	m := new(HashMap[String, int]).Init(5)
	m.Insert("a", 5)
	m.Insert("b", 3)
	m.Insert("c", 55)
	m.Insert("d", 100)
	m.Insert("e", 299)
	m.Insert("f", 444)
	m.Set("a", 150)
	m.Insert("a", 233)
	m.Pop("e")
	m.Print()
}
