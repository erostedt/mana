package main

import (
    "fmt"
    "os"
    "mana/pkg/mana"
)

func usage() {
    fmt.Println("Usage:")
    fmt.Println("    go run ./cmd/edit_distance/main.go <string1> <string2>")
}

func main() {
    if len(os.Args) != 3 {
        fmt.Printf("Invalid number of arguments: %d\n", len(os.Args) - 1)
        usage()
        return
    }

    distance := mana.EditDistance([]rune(os.Args[1]), []rune(os.Args[2]))
    fmt.Printf("Edit distance: %d\n", distance)
}
