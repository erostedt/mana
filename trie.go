package main

import (
	"fmt"
	"strings"
)

type Ascii byte

func (ascii Ascii) Hash() uint {
	return uint(ascii)
}

type TrieNode struct {
	ascii      Ascii
	children   HashMap[Ascii, *TrieNode]
	isTerminal bool
}

type Trie struct {
	root *TrieNode
	cap  uint
}

func NewNode(ascii Ascii, isTerminal bool, cap uint) *TrieNode {
	node := new(TrieNode)
	node.ascii = ascii
	node.isTerminal = isTerminal
	node.children = *new(HashMap[Ascii, *TrieNode]).Init(cap)
	return node
}

func (t *Trie) Init(cap uint) *Trie {
	t.root = NewNode(0, false, cap)
	t.cap = cap
	return t
}

func (t *Trie) Insert(word []Ascii) {
	node := t.root
	for _, ascii := range word {
		n, err := node.children.Get(ascii)
		if err == nil {
			node = n
		} else {
			newNode := NewNode(ascii, false, t.cap)
			node.children.Insert(ascii, newNode)
			node = newNode
		}
	}
	node.isTerminal = true
}

func (t *Trie) Contains(word []Ascii) bool {
	node, err := t.FindTail(word)
	return err == nil && node.isTerminal
}

func (t *Trie) StartsWith(word []Ascii) bool {
	_, err := t.FindTail(word)
	return err == nil
}

func (t *Trie) PrintSuggestions(word []Ascii) {
	tail, err := t.FindTail(word)
	if err == nil {
		tail.RecPrintSuggestion(word)
	}
}

func (t *TrieNode) RecPrintSuggestion(base []Ascii) {
	iter := t.children.CreateIterator()
	for iter.HasNext() {
		node := iter.Next()
		word := strings.Join([]string{string(base), string(node.value.ascii)}, "")
		if node.value.isTerminal {
			fmt.Println(word)
		}
		node.value.RecPrintSuggestion([]Ascii(word))
	}
}

func (t *Trie) FindTail(word []Ascii) (*TrieNode, error) {
	node := t.root
	for _, ascii := range word {
		n, err := node.children.Get(ascii)
		if err != nil {
			return n, err
		}
		node = n
	}
	return node, nil
}

func (t *Trie) Print() {
	t.root.RecPrint()
}

func (t *TrieNode) RecPrint() {
	iterator := t.children.CreateIterator()
	for iterator.HasNext() {
		bucket := iterator.Next()
		fmt.Println(string(bucket.key))
		bucket.value.RecPrint()
	}
}
