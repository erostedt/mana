package main

import (
	"errors"
	"fmt"
)

type String string

type ForwardIterable[T any] interface {
	HasNext() bool
	Next() *T
}

type Hashable interface {
	Hash() uint
}

type Key interface {
	Hashable
	comparable
}

type Bucket[K Key, V any] struct {
	key      K
	value    V
	occupied bool
}

type HashMap[K Key, V any] struct {
	buckets     []Bucket[K, V]
	size        uint
	cap         uint
	extendLimit uint
}

type HashMapIterator[key Key, value any] struct {
	buckets           []Bucket[key, value]
	currentIndex      uint
	numVisitedBuckets uint
	numFilledBuckets  uint
}

func (m *HashMap[Key, Value]) Init(cap uint) *HashMap[Key, Value] {
	m.size = 0
	m.extendLimit = (cap / 3) * 2
	m.cap = cap

	m.buckets = make([]Bucket[Key, Value], cap)
	return m
}

func (m *HashMap[Key, Value]) Insert(key Key, value Value) bool {
	if m.size >= m.extendLimit {
		m.Extend()
	}
	return m._Insert(key, value)
}

func (m *HashMap[Key, Value]) Extend() {
	buckets := m.buckets

	m.size = 0
	m.cap = (m.cap + 1) * 2
	m.extendLimit = (m.cap / 3) * 2
	m.buckets = make([]Bucket[Key, Value], m.cap)

	for _, bucket := range buckets {
		if bucket.occupied {
			m._Insert(bucket.key, bucket.value)
		}
	}
}

func (m *HashMap[Key, Value]) _Insert(key Key, value Value) bool {
	slot := key.Hash() % m.cap
	var i uint = 0
	for ; i < m.cap; i++ {
		if m.buckets[slot].key == key {
			return false
		}
		if !m.buckets[slot].occupied {
			m.buckets[slot] = Bucket[Key, Value]{key, value, true}
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
		return m.buckets[slot].value, nil
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
		m.buckets[slot].value = value
		return nil
	}
	return err
}

func (m *HashMap[Key, Value]) SetInsert(key Key, value Value) {
	err := m.Set(key, value)
	if err != nil {
		m.Insert(key, value)
	}
}

func (m *HashMap[Key, Value]) Pop(key Key) (Value, error) {
	slot, err := m._FindSlot(key)
	if err == nil {
		value := m.buckets[slot].value
		m.buckets[slot].occupied = false
		m.size--
		return value, nil
	}
	return *new(Value), errors.New("could not pop item, since it was not present")
}

func (m *HashMap[Key, Value]) Contains(key Key) bool {
	_, err := m._FindSlot(key)
	return err == nil
}

func (m *HashMap[Key, Value]) _FindSlot(key Key) (uint, error) {
	hash := key.Hash()
	slot := hash % m.cap
	for {
		if !m.buckets[slot].occupied {
			return 0, errors.New("key not found")
		}

		if m.buckets[slot].key == key {
			return slot, nil
		}

		slot = (slot*5 + 1 + hash) % m.cap
		hash >>= 5
	}
}

func (m *HashMap[Key, Value]) Print() {
	for _, bucket := range m.buckets {
		if bucket.occupied {
			fmt.Printf("%+v: %+v \n", bucket.key, bucket.value)
		}
	}
}

func (m *HashMap[Key, Value]) FullPrint() {
	for _, bucket := range m.buckets {
		if bucket.occupied {
			fmt.Printf("%+v: %+v \n", bucket.key, bucket.value)
		} else {
			fmt.Println("[X]")
		}
	}
}

func (m *HashMap[Key, Value]) CreateIterator() HashMapIterator[Key, Value] {
	return HashMapIterator[Key, Value]{
		m.buckets,
		0,
		0,
		m.size,
	}
}

func (i *HashMapIterator[Key, Value]) HasNext() bool {
	return i.numVisitedBuckets < i.numFilledBuckets
}

func (i *HashMapIterator[Key, Value]) Next() *Bucket[Key, Value] {
	if i.HasNext() {
		for !i.buckets[i.currentIndex].occupied {
			i.currentIndex++
		}
		bucket := &i.buckets[i.currentIndex]
		i.currentIndex++
		i.numVisitedBuckets++
		return bucket
	}
	return nil
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
}
