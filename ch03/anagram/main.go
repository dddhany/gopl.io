package main

import (
	"fmt"
)

func anagram(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("%v\n", anagram("abcdcba"))
	fmt.Printf("%v\n", anagram("abcdecba"))
}
