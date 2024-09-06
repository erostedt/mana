package mana

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

type BucketState int

const (
	BucketVacant BucketState = iota
	BucketOccupied
	BucketTombstone
)

type Bucket[K Key, V any] struct {
	key   K
	value V
	state BucketState
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

func NewDict[K Key, V any]() *Dict[K, V] {
	d := MakeDict[K, V]()
	return &d
}

func MakeDict[K Key, V any]() Dict[K, V] {
    const INITIAL_CAPACITY uint = 8
	return Dict[K, V]{
		buckets:     make([]Bucket[K, V], INITIAL_CAPACITY),
		size:        0,
		cap:         INITIAL_CAPACITY,
		extendLimit: (INITIAL_CAPACITY / 3) * 2,
	}
}

func (d *Dict[K, V]) Insert(key K, value V) bool {
	if d.size >= d.extendLimit {
		d.Extend()
	}
	return d.insert(key, value)
}

func (d *Dict[K, V]) Extend() {
	buckets := d.buckets

	d.size = 0
	d.cap = (d.cap + 1) * 2
	d.extendLimit = (d.cap / 3) * 2
	d.buckets = make([]Bucket[K, V], d.cap)

	for _, bucket := range buckets {
		if bucket.state == BucketOccupied {
			d.insert(bucket.key, bucket.value)
		}
	}
}

func (d *Dict[K, V]) insert(key K, value V) bool {
	slot := key.Hash() % d.cap
	var i uint = 0
	for ; i < d.cap; i++ {
		if d.buckets[slot].key == key {
			return false
		}
		if d.buckets[slot].state != BucketOccupied {
			d.buckets[slot] = Bucket[K, V]{key, value, BucketOccupied}
			d.size++
			return true
		}
		slot = (slot + 1) % d.cap
	}
	return false
}

func (d *Dict[K, V]) Get(key K) (V, error) {
	slot, err := d.findSlot(key)
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
	slot, err := d.findSlot(key)
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

func (d *Dict[K, V]) Contains(key K) bool {
	_, err := d.findSlot(key)
	return err == nil
}

func (d *Dict[K, V]) Pop(key K) (V, error) {
	slot, err := d.findSlot(key)
	if err != nil {
		return *new(V), err
	}
	value := d.buckets[slot].value
	d.buckets[slot].state = BucketTombstone
	return value, nil
}

func (d *Dict[K, V]) findSlot(key K) (uint, error) {
	hash := key.Hash()
	slot := hash % d.cap
	for {
		bucket := d.buckets[slot]
		if bucket.state == BucketVacant {
			return 0, errors.New("key not found")
		}

		if bucket.state == BucketOccupied && bucket.key == key {
			return slot, nil
		}

		slot = (slot*5 + 1 + hash) % d.cap
		hash >>= 5
	}
}

func (d *Dict[K, V]) Print() {
	for _, bucket := range d.buckets {
		if bucket.state == BucketOccupied {
			fmt.Printf("%+v: %+v \n", bucket.key, bucket.value)
		}
	}
}

func (d *Dict[K, V]) FullPrint() {
	for _, bucket := range d.buckets {
		switch bucket.state {
		case BucketOccupied:
			fmt.Printf("%+v: %+v \n", bucket.key, bucket.value)
		case BucketVacant:
			fmt.Println("[ ]")

		case BucketTombstone:
			fmt.Println("[T]")
		default:
			panic("Unexpected BucketState")
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
	if !i.HasNext() {
        return nil;
    }
    for i.buckets[i.currentIndex].state != BucketOccupied {
        i.currentIndex++
    }
    bucket := &i.buckets[i.currentIndex]
    i.currentIndex++
    i.numVisitedBuckets++
    return bucket
}
