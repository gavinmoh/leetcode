package main

import "strings"

func addSpaces(s string, spaces []int) string {
	var result strings.Builder
	spaceIdx := 0
	n := len(spaces)
	for i, char := range s {
		if spaceIdx < n && i == spaces[spaceIdx] {
			result.WriteRune(' ')
			spaceIdx++
		}
		result.WriteRune(char)
	}

	return result.String()
}
