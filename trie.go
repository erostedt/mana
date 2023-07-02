package main

import (
	"encoding/binary"
	"fmt"
	"strings"
	"unicode"
)

type Rune rune

func djb2(bytes []byte) uint {
	var hash uint = 5381

	for _, c := range bytes {
		hash = (((hash << 5) + hash) + uint(c))
	}
	return hash
}

func (char Rune) Hash() uint {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(char))
	return djb2(bytes)
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

func newNode(char Rune, parent *TrieNode, isTerminal bool, cap uint) *TrieNode {
	node := new(TrieNode)
	node.parent = parent
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
	return Trie{root: newNode(0, nil, false, cap), cap: cap}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		lowercased := Rune(unicode.ToLower(char))
		n, err := node.children.Get(lowercased)
		if err == nil {
			node = n
		} else {
			newNode := newNode(lowercased, node, false, t.cap)
			node.children.Insert(lowercased, newNode)
			node = newNode
		}
	}
	node.isTerminal = true
}

func (t *Trie) Contains(word string) bool {
	node, err := t.findTail(word)
	return err == nil && node.isTerminal
}

func (t *Trie) StartsWith(word string) bool {
	_, err := t.findTail(word)
	return err == nil
}

func (t *Trie) Autocomplete(base string, maxSuggestions int) []string {
	tail, err := t.findTail(base)

	if err == nil {
		words := tail.bfs(maxSuggestions)
		var caseFunc func(string) string
		if isCapitalCase(base) {
			caseFunc = strings.ToUpper
		} else if firstIsUpper(base) {
			caseFunc = toTitleCase
		} else {
			caseFunc = strings.ToLower
		}

		for i := 0; i < len(words); i++ {
			words[i] = caseFunc(base + words[i])
		}

		return words
	}
	return make([]string, 0)
}

func (root *TrieNode) bfs(maxLimit int) []string {
	queue := MakeDeque[*TrieNode]()
	queue.AddLast(root)
	words := make([]string, 0, maxLimit)

	for len(words) < maxLimit {
		node, err := queue.PopFirst()
		if err != nil {
			break
		}

		if node.isTerminal {
			words = append(words, node.backtrack(root))
		}

		iter := node.children.CreateIterator()
		for iter.HasNext() {
			bucket := iter.Next()
			queue.AddLast(bucket.value)
		}
	}
	return words
}

func (t *TrieNode) backtrack(tail *TrieNode) string {
	word := make([]rune, 0)
	node := t
	for node != tail {
		word = append(word, rune(node.char))
		node = node.parent
	}
	reverse[rune](word)
	return string(word)
}

func reverse[T any](slice []T) {
	length := len(slice)
	for forwardIndex := 0; forwardIndex < length/2; forwardIndex++ {
		backwardIndex := length - forwardIndex - 1
		tmp := slice[forwardIndex]
		slice[forwardIndex] = slice[backwardIndex]
		slice[backwardIndex] = tmp
	}
}

func (t *Trie) findTail(word string) (*TrieNode, error) {
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
	t.root.recPrint()
}

func (t *TrieNode) recPrint() {
	iterator := t.children.CreateIterator()
	for iterator.HasNext() {
		bucket := iterator.Next()
		fmt.Println(string(bucket.key))
		bucket.value.recPrint()
	}
}

func firstIsUpper(s string) bool {
	if len(s) == 0 {
		return false
	}

	chars := []rune(s)

	return unicode.IsUpper(chars[0])
}

func isCapitalCase(s string) bool {
	if len(s) < 2 {
		return false
	}

	chars := []rune(s)
	for i := 0; i < len(chars); i++ {
		if unicode.IsLower(chars[i]) {
			return false
		}
	}
	return true
}

func toTitleCase(s string) string {
	if len(s) == 0 {
		return s
	}
	chars := []rune(strings.ToLower(s))
	chars[0] = unicode.ToUpper(chars[0])
	return string(chars)
}
