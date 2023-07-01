package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("animals.txt")

	if err != nil {
		fmt.Print("Could not read file.")
		os.Exit(1)
	}

	trie := NewTrie(5)
	animals := strings.Split(string(data), "\n")
	for _, animal := range animals {
		trie.Insert(animal)
	}

	fmt.Println(trie.Autocomplete("Pe", 3))
}
