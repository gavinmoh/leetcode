package main

import "strings"

func canBeTypedWords(text string, brokenLetters string) int {
	broken := make(map[rune]bool)
	for _, c := range brokenLetters {
		broken[c] = true
	}

	count := 0
	for _, word := range strings.Split(text, " ") {
		canType := true
		for _, c := range word {
			if broken[c] {
				canType = false
				break
			}
		}

		if canType {
			count++
		}
	}

	return count
}
