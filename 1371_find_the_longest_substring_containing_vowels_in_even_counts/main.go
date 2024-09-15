package main

func findTheLongestSubstring(s string) int {
	vowels := map[rune]int{'a': 1, 'e': 2, 'i': 4, 'o': 8, 'u': 16}
	prefixXOR := make([]int, len(s))
	currentXOR := 0
	length := 0
	firstOccurrence := make([]int, 32)
	for i := 0; i < 32; i++ {
		firstOccurrence[i] = -1
	}

	for i, char := range s {
		if value, ok := vowels[char]; ok {
			currentXOR = currentXOR ^ value
		}
		prefixXOR[i] = currentXOR
		if firstOccurrence[currentXOR] == -1 && currentXOR != 0 {
			firstOccurrence[currentXOR] = i
		}
		length = max(length, i-firstOccurrence[currentXOR])
	}

	return length
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
