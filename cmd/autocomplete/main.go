package main

import (
	"bufio"
	"fmt"
	"mana/pkg/mana"
	"os"
	"strconv"
	"strings"
)

func usage() {
    fmt.Println("Usage:")
    fmt.Println("    add <word1> <word2> ...")
    fmt.Println("    suggest <subword1> <subword2> ...")
    fmt.Println("    max_suggestions <count>")
    fmt.Println("    exit")
    fmt.Println("    help")
}

func parseArgument(reader *bufio.Reader) (command string, args []string) {
    input, err := reader.ReadString('\n')
    if err != nil {
        panic(err)
    }
    splitted := strings.Split(strings.TrimSpace(input), " ")
    if (len(splitted) == 0) {
        return command, args
    }
    command = splitted[0]

    if (len(splitted) > 1) {
        args = splitted[1:]
    }
    return command, args

}

func addWords(trie *mana.Trie, words []string) {
    if (len(words) == 0) {
        fmt.Println("Invalid usage.")
        usage()
    }
    for _, word := range words {
        trie.Insert(word)
    }
}

func suggest(trie *mana.Trie, words []string, maxSuggestions int) {
    if (len(words) == 0) {
        fmt.Println("Invalid usage.")
        usage()
    }
    for _, word := range words {
        fmt.Printf("Suggestions for: %s\n", word)
        suggestions := trie.Autocomplete(word,  maxSuggestions)
        for _, suggestion := range suggestions {
            fmt.Printf("  - %s\n", suggestion)
        }
        fmt.Println("")
    }
}

func setMaxSuggestions(args []string, maxSuggestions *int) {
    if (len(args) == 0) {
        fmt.Println("Invalid usage.")
        usage()
    }

    integer, err := strconv.ParseInt(args[0], 10, 64)
    if err != nil {
        fmt.Println("Positive integer not provided")
        usage()
        return
    }

    if (integer <= 0) {
        fmt.Println("Integer must be larger than 0")
        usage()
        return
    }
    *maxSuggestions = int(integer)
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    trie := mana.MakeTrie()
    maxSuggestions := 3

    usage()
    for {
        command, args := parseArgument(reader)
        switch command {
            case "add":
                addWords(&trie, args)
            case "suggest":
                suggest(&trie, args, maxSuggestions)
            case "max_suggestions":
                setMaxSuggestions(args, &maxSuggestions)
            case "exit":
                os.Exit(0)
            case "help":
                usage()
            case "":
                fmt.Println("Missing command")
                usage()
            default:
                fmt.Printf("Invalid command: %s\n", command)
                usage()
            }
    }
}
