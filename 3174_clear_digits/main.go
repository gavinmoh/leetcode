package main

import "strings"

func clearDigits(s string) string {
	marked := make([]bool, len(s))

	for i, c := range s {
		if isDigit(c) {
			marked[i] = true

			for left := i; left >= 0; left-- {
				if !marked[left] && !isDigit(rune(s[left])) {
					marked[left] = true
					break
				}
			}
		}
	}

	var result strings.Builder
	for i, c := range s {
		if !marked[i] {
			result.WriteRune(c)
		}
	}

	return result.String()
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}
