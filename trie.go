package main

import (
	"fmt"
	"strings"
)

type Rune rune

func (char Rune) Hash() uint {
	// TODO: FIXME
	return uint(char)
}

type TrieNode struct {
	char       Rune
	children   Dict[Rune, *TrieNode]
	isTerminal bool
}

type Trie struct {
	root *TrieNode
	cap  uint
}

func NewNode(char Rune, isTerminal bool, cap uint) *TrieNode {
	node := new(TrieNode)
	node.char = char
	node.isTerminal = isTerminal
	node.children = MakeDict[Rune, *TrieNode](cap)
	return node
}

func NewTrie(cap uint) *Trie {
	t := MakeTrie(cap)
	return &t
}

func MakeTrie(cap uint) Trie {
	return Trie{root: NewNode(0, false, cap), cap: cap}
}

func (t *Trie) Insert(word []Rune) {
	node := t.root
	for _, char := range word {
		n, err := node.children.Get(char)
		if err == nil {
			node = n
		} else {
			newNode := NewNode(char, false, t.cap)
			node.children.Insert(char, newNode)
			node = newNode
		}
	}
	node.isTerminal = true
}

func (t *Trie) Contains(word []Rune) bool {
	node, err := t.FindTail(word)
	return err == nil && node.isTerminal
}

func (t *Trie) StartsWith(word []Rune) bool {
	_, err := t.FindTail(word)
	return err == nil
}

func (t *Trie) PrintSuggestions(word []Rune) {
	tail, err := t.FindTail(word)
	if err == nil {
		tail.RecPrintSuggestion(word)
	}
}

func (t *TrieNode) RecPrintSuggestion(base []Rune) {
	iter := t.children.CreateIterator()
	for iter.HasNext() {
		node := iter.Next()
		word := strings.Join([]string{string(base), string(node.value.char)}, "")
		if node.value.isTerminal {
			fmt.Println(word)
		}
		node.value.RecPrintSuggestion([]Rune(word))
	}
}

func (t *Trie) FindTail(word []Rune) (*TrieNode, error) {
	node := t.root
	for _, char := range word {
		n, err := node.children.Get(char)
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
