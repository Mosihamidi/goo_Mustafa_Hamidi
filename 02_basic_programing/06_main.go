package main

import (
	"fmt"
	"strings"
)

func main() {
	word1 := "cafe"
	word2 := "face"

	if isAnagram(word1, word2) {
		fmt.Println("Anagram")
	} else {
		fmt.Println("Not Anagram")
	}
}

func isAnagram(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	word1 = strings.ToLower(word1)
	word2 = strings.ToLower(word2)

	for i := 0; i < len(word1); i++ {
		if strings.Count(word1, string(word1[i])) != strings.Count(word2, string(word1[i])) {
			return false
		}
	}

	return true
}
