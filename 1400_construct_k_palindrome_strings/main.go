package main

func canConstruct(s string, k int) bool {
	if k > len(s) {
		return false
	}

	charFreq := map[rune]int{}
	for _, char := range s {
		charFreq[char]++
	}

	oddCount := 0
	for _, freq := range charFreq {
		if freq%2 != 0 {
			oddCount++
		}
	}

	return oddCount <= k
}
