package main

import (
	"fmt"
)

func main() {
	words := []string{"a", "a", "a", "b", "b", "c", "d", "e", "e"}
	words = removeAdjacentDuplicate(words)
	fmt.Println(words)
}

func removeAdjacentDuplicate(s []string) []string {
	if len(s) <= 1 {
		return s
	}

	var p int
	for i := 1; i < len(s); i++ {
		if s[p] == s[i] {
			continue
		}

		p++
		s[p] = s[i]
	}

	return s[:p+1]
}
