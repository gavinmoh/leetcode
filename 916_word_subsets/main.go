package main

func wordSubsets(words1 []string, words2 []string) []string {
	charMaxFreq := map[rune]int{}
	for _, word := range words2 {
		charCount := map[rune]int{}
		for _, char := range word {
			charCount[char]++
		}

		for char, count := range charCount {
			if maxFreq, ok := charMaxFreq[char]; ok {
				charMaxFreq[char] = max(maxFreq, count)
			} else {
				charMaxFreq[char] = count
			}
		}
	}

	result := []string{}
	for _, word := range words1 {
		charCount := map[rune]int{}
		for _, char := range word {
			charCount[char]++
		}

		isUniversal := true
		for char, maxFreq := range charMaxFreq {
			if count, ok := charCount[char]; ok && count >= maxFreq {
				continue
			}

			isUniversal = false
			break
		}

		if isUniversal {
			result = append(result, word)
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
