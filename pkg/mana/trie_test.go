package mana

import (
	"testing"
)


var animals []string = []string {
    "Lion",
    "Tiger",
    "Elephant",
    "Giraffe",
    "Monkey",
    "Zebra",
    "Kangaroo",
    "Hippopotamus",
    "Koala",
    "Panda",
    "Gorilla",
    "Crocodile",
    "Dolphin",
    "Ostrich",
    "Peacock",
    "Sloth",
    "Raccoon",
    "Penguin",
    "Cheetah",
    "Rhino",
    "Koala",
    "Jaguar",
    "Platypus",
    "PolarBear",
}

func TestTrieComplete(t *testing.T) {
	trie := NewTrie()
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
