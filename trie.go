package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Rune rune

func (char Rune) Hash() uint {
	// TODO: FIXME
	return uint(char)
}

type TrieNode struct {
	char       Rune
	parent     *TrieNode
	children   Dict[Rune, *TrieNode]
	isTerminal bool
}

type Trie struct {
	root *TrieNode
	cap  uint
}

func NewNode(char Rune, parent *TrieNode, isTerminal bool, cap uint) *TrieNode {
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
	return Trie{root: NewNode(0, nil, false, cap), cap: cap}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		lowercased := Rune(unicode.ToLower(char))
		n, err := node.children.Get(lowercased)
		if err == nil {
			node = n
		} else {
			newNode := NewNode(lowercased, node, false, t.cap)
			node.children.Insert(lowercased, newNode)
			node = newNode
		}
	}
	node.isTerminal = true
}

func (t *Trie) Contains(word string) bool {
	node, err := t.FindTail(word)
	return err == nil && node.isTerminal
}

func (t *Trie) StartsWith(word string) bool {
	_, err := t.FindTail(word)
	return err == nil
}

func (t *Trie) PrintSuggestions(word string) {
	tail, err := t.FindTail(word)
	if err == nil {
		tail.RecPrintSuggestion(word)
	}
}

func (t *TrieNode) RecPrintSuggestion(base string) {
	iter := t.children.CreateIterator()
	for iter.HasNext() {
		node := iter.Next()
		word := strings.Join([]string{base, string(node.value.char)}, "")
		if node.value.isTerminal {
			fmt.Println(word)
		}
		node.value.RecPrintSuggestion(word)
	}
}

func (t *Trie) Autocomplete(base string) []string {
	// Finish me
	maxSuggestions := 3
	options := make([]string, 0, maxSuggestions)
	tail, err := t.FindTail(base)

	if err != nil {
		if tail.isTerminal {
			options = append(options, tail.Backtrack())
		}
		tailOptions := tail.bfs(maxSuggestions - len(options))
		options = append(options, tailOptions...)
	}

	//BackTrack
	return options
}

func (root *TrieNode) bfs(maxLimit int) []string {
	queue := MakeDeque[*TrieNode]()
	queue.AddLast(root)
	words := make([]string, 0)

	for len(words) < maxLimit {
		node, err := queue.PopFirst()
		if err != nil {
			break
		}

		if node.isTerminal {
			words = append(words, node.Backtrack())
		}

		iter := node.children.CreateIterator()
		for iter.HasNext() {
			bucket := iter.Next()
			queue.AddLast(bucket.value)
		}
	}

	return words
}

func (t *TrieNode) Backtrack() string {
	word := make([]rune, 0)
	node := t
	for node != nil {
		word = append(word, rune(node.char))
		node = node.parent
	}
	Reverse[rune](word)
	return string(word)
}

func Reverse[T any](slice []T) {
	length := len(slice)
	for forwardIndex := 0; forwardIndex < length/2; forwardIndex++ {
		backwardIndex := length - forwardIndex - 1
		tmp := slice[forwardIndex]
		slice[forwardIndex] = slice[backwardIndex]
		slice[backwardIndex] = tmp
	}
}

func (t *Trie) FindTail(word string) (*TrieNode, error) {
	node := t.root
	for _, char := range word {
		lowercased := Rune(unicode.ToLower(char))
		n, err := node.children.Get(lowercased)
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
