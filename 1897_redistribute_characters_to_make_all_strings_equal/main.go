package main

func makeEqual(words []string) bool {
	n := len(words)
	freqMap := make(map[rune]int)

	for _, word := range words {
		for _, char := range word {
			if _, ok := freqMap[char]; ok {
				freqMap[char]++
			} else {
				freqMap[char] = 1
			}
		}
	}

	for _, freq := range freqMap {
		if freq%n != 0 {
			return false
		}
	}

	return true
}
