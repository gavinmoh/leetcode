package main

import "strings"

func findWordsContaining(words []string, x byte) []int {
	result := []int{}
	for i, word := range words {
		if strings.ContainsAny(word, string(x)) {
			result = append(result, i)
		}
	}

	return result
}
