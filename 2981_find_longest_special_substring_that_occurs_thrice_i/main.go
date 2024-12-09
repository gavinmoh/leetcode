package main

import "strings"

func maximumLength(s string) int {
	substrMap := map[string]int{}

	for left := 0; left < len(s); left++ {
		var b strings.Builder
		for right := left; right < len(s); right++ {
			if right != left && s[right] != s[left] {
				break
			}
			b.WriteByte(s[right])
			substrMap[b.String()]++
		}
	}

	maxLength := -1
	for substr, count := range substrMap {
		if count < 3 {
			continue
		}
		if len(substr) > maxLength {
			maxLength = len(substr)
		}
	}

	return maxLength
}
