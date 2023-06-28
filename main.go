package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Match struct {
	distance int
	word     string
}

func main() {

	dictionary := [][]rune{
		[]rune("hello"),
		[]rune("yellow"),
		[]rune("testtest"),
	}

	queryWord := []rune("hello")

	closestWords := []Match{}
	limit := 3
	for _, word := range dictionary {

		dp, m, n := levenstein(word, queryWord)
		if dp[m*n-1] > limit {
			continue
		}
		closestWords = append(closestWords, Match{dp[m*n-1], string(word)})
	}

	if len(closestWords) > 0 {
		sort.SliceStable(closestWords, func(i, j int) bool { return closestWords[i].distance < closestWords[j].distance })
		fmt.Println("Did you mean: ")
	}

	for _, match := range closestWords {
		fmt.Println(match.word)
	}
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
	//trie.Print()
	ABC()
}
