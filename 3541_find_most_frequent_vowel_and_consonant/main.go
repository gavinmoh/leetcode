package main

func maxFreqSum(s string) int {
	freq := make([]int, 26)
	for _, c := range s {
		freq[int(c-'a')]++
	}

	maxVowel, maxConsonant := 0, 0
	for i, count := range freq {
		switch rune('a' + i) {
		case 'a', 'e', 'i', 'o', 'u':
			maxVowel = max(maxVowel, count)
		default:
			maxConsonant = max(maxConsonant, count)
		}
	}

	return maxVowel + maxConsonant
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
