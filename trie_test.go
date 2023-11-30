package main

import (
	"os"
	"strings"
	"testing"
)

func TestTrieComplete(t *testing.T) {
	data, err := os.ReadFile("animals.txt")
	if err != nil {
		t.Error("Could not read file.")
	}

	trie := NewTrie(5)
	animals := strings.Split(string(data), "\n")
	for _, animal := range animals {
		trie.Insert(animal)
	}

	suggestions := trie.Autocomplete("Pe", 2)

	if suggestions[0] != "Penguin" {
		t.Errorf("First suggestion should have been Penguin, was %s", suggestions[0])
	}

	if suggestions[1] != "Peacock" {
		t.Errorf("Second suggestion should have been Peacock, was %s", suggestions[1])
	}
}
