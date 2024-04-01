package main

import "strings"

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)
	last_word := words[len(words)-1]

	return len(last_word)
}
