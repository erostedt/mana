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

	dictionary := [][]byte{
		[]byte("hello"),
		[]byte("yellow"),
	}

	queryWord := "hello"
	asciiSlice := []byte(queryWord)

	closestWords := []Match{}
	limit := 3
	for _, word := range dictionary {

		dp, m, n := levenstein(word, asciiSlice)
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
	a()
	data, err := os.ReadFile("animals.txt")

	if err != nil {
		fmt.Print("Could not read file.")
		os.Exit(1)
	}

	trie := new(Trie).Init(5)
	animals := strings.Split(string(data), "\n")
	for _, animal := range animals {
		trie.Insert([]Ascii(animal))
	}
	trie.PrintSuggestions([]Ascii("Pe"))
	//trie.Print()

}
