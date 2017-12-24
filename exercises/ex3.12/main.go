package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(anagram("banana", "Banana"))
}

func anagram(word1, word2 string) bool {
	word1 = strings.ToLower(strings.TrimSpace(word1))
	word2 = strings.ToLower(strings.TrimSpace(word2))

	if word1 == word2 {
		return false
	}

	runes1 := []rune(word1)
	runes2 := []rune(word2)

	sort.Slice(runes1, func(i int, j int) bool {
		return runes1[i] < runes1[j]
	})

	sort.Slice(runes2, func(i int, j int) bool {
		return runes2[i] < runes2[j]
	})

	return string(runes1) == string(runes2)
}
