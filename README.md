mana is a small project regarding edit_distances/stringdiff and autocomplete.

This project is not intended for any professional use, but it is free to use by whomever wants to.

The main package is in pkg/mana and contains:
 - deque.go -> Implementation of a double ended queue. (This comes shipped with go as list.List, but I wanted to make my own)
 - dict.go  -> Implementation of a dictionary/hashmap. (This comes shipped with go as map      , but I wanted to make my own)
 - edit_distance.go -> Implementation of levenstein edit_distance, useful for comparing how similar two strings are.
 - trie.go -> Implementation of the `trie` datastructure which can be used for autocomplete.
 - corresponding tests.

This project also comes with two programs: 
 - cmd/autocomplete/main.go -> Sample program of how a trie can be used for autocomplete.
 - cmd/edit_distance/main.go -> Sample program for comparing two strings using edit_distance
