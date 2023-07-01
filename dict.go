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

type Dict[K Key, V any] struct {
	buckets     []Bucket[K, V]
	size        uint
	cap         uint
	extendLimit uint
}

type DictIterator[K Key, V any] struct {
	buckets           []Bucket[K, V]
	currentIndex      uint
	numVisitedBuckets uint
	numFilledBuckets  uint
}

func NewDict[K Key, V any](cap uint) *Dict[K, V] {
	d := MakeDict[K, V](cap)
	return &d
}

func MakeDict[K Key, V any](cap uint) Dict[K, V] {
	return Dict[K, V]{
		buckets:     make([]Bucket[K, V], cap),
		size:        0,
		cap:         cap,
		extendLimit: (cap / 3) * 2,
	}
}

func (d *Dict[K, V]) Insert(key K, value V) bool {
	if d.size >= d.extendLimit {
		d.Extend()
	}
	return d._Insert(key, value)
}

func (d *Dict[K, V]) Extend() {
	buckets := d.buckets

	d.size = 0
	d.cap = (d.cap + 1) * 2
	d.extendLimit = (d.cap / 3) * 2
	d.buckets = make([]Bucket[K, V], d.cap)

	for _, bucket := range buckets {
		if bucket.occupied {
			d._Insert(bucket.key, bucket.value)
		}
	}
}

func (d *Dict[K, V]) _Insert(key K, value V) bool {
	slot := key.Hash() % d.cap
	var i uint = 0
	for ; i < d.cap; i++ {
		if d.buckets[slot].key == key {
			return false
		}
		if !d.buckets[slot].occupied {
			d.buckets[slot] = Bucket[K, V]{key, value, true}
			d.size++
			return true
		}
		slot = (slot + 1) % d.cap
	}
	return false
}

func (d *Dict[K, V]) Get(key K) (V, error) {
	slot, err := d._FindSlot(key)
	if err == nil {
		return d.buckets[slot].value, nil
	}
	return *new(V), errors.New("key not found")
}

func (d *Dict[K, V]) GetDefault(key K, def V) V {
	value, err := d.Get(key)
	if err == nil {
		return value
	}
	return def
}

func (d *Dict[K, V]) Set(key K, value V) error {
	slot, err := d._FindSlot(key)
	if err == nil {
		d.buckets[slot].value = value
		return nil
	}
	return err
}

func (d *Dict[K, V]) SetInsert(key K, value V) {
	err := d.Set(key, value)
	if err != nil {
		d.Insert(key, value)
	}
}

func (d *Dict[K, V]) Pop(key K) (V, error) {
	slot, err := d._FindSlot(key)
	if err == nil {
		value := d.buckets[slot].value
		d.buckets[slot].occupied = false
		d.size--
		return value, nil
	}
	return *new(V), errors.New("could not pop item, since it was not present")
}

func (d *Dict[K, V]) Contains(key K) bool {
	_, err := d._FindSlot(key)
	return err == nil
}

func (d *Dict[K, V]) _FindSlot(key K) (uint, error) {
	hash := key.Hash()
	slot := hash % d.cap
	for {
		if !d.buckets[slot].occupied {
			return 0, errors.New("key not found")
		}

		if d.buckets[slot].key == key {
			return slot, nil
		}

		slot = (slot*5 + 1 + hash) % d.cap
		hash >>= 5
	}
}

func (d *Dict[K, V]) Print() {
	for _, bucket := range d.buckets {
		if bucket.occupied {
			fmt.Printf("%+v: %+v \n", bucket.key, bucket.value)
		}
	}
}

func (d *Dict[K, V]) FullPrint() {
	for _, bucket := range d.buckets {
		if bucket.occupied {
			fmt.Printf("%+v: %+v \n", bucket.key, bucket.value)
		} else {
			fmt.Println("[X]")
		}
	}
}

func (d *Dict[K, V]) CreateIterator() DictIterator[K, V] {
	return DictIterator[K, V]{
		d.buckets,
		0,
		0,
		d.size,
	}
}

func (i *DictIterator[K, V]) HasNext() bool {
	return i.numVisitedBuckets < i.numFilledBuckets
}

func (i *DictIterator[K, V]) Next() *Bucket[K, V] {
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
