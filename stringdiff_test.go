package main

import (
	"sort"
	"testing"
)

func TestStringDiff(t *testing.T) {
	words := [][]rune{
		[]rune("hello"),
		[]rune("yellow"),
		[]rune("testtest"),
	}

	queryWord := []rune("helloo")

	closestWords := []Match{}
	limit := 3
	for _, word := range words {

		distance := levenstein(word, queryWord)
		if distance > limit {
			continue
		}
		closestWords = append(closestWords, Match{string(word), distance})
	}

	sort.SliceStable(closestWords, func(i, j int) bool { return closestWords[i].distance < closestWords[j].distance })

	if len(closestWords) != 2 {
		t.Errorf("Invalid number of close words, got %d, expected %d", len(closestWords), 2)
	}
}
