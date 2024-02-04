package main

import "math"

func minWindow(s string, t string) string {
	m, n := len(s), len(t)

	if n > m {
		return ""
	}

	tMap := make(map[rune]int)
	for _, char := range t {
		tMap[char]++
	}

	left, sLeft := 0, -1
	charContained := 0
	minSize := math.MaxInt

	windowMap := make(map[rune]int)
	for index, char := range s {
		windowMap[char]++

		if freq, ok := tMap[char]; ok && freq >= windowMap[char] {
			charContained++
		}

		for charContained == n {
			size := index - left + 1
			if minSize > size {
				minSize = size
				sLeft = left
			}
			lChar := rune(s[left])
			if freq, ok := tMap[lChar]; ok && freq >= windowMap[lChar] {
				charContained--
			}
			windowMap[lChar]--
			left++
		}
	}

	if sLeft == -1 {
		return ""
	}

	return s[sLeft : sLeft+minSize]
}
