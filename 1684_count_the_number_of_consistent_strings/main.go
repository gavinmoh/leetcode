package main

func countConsistentStrings(allowed string, words []string) int {
	hash := make([]bool, 26)
	for _, letter := range allowed {
		hash[letter-'a'] = true
	}

	count := 0
	for _, word := range words {
		consistent := true
		for _, letter := range word {
			if !hash[letter-'a'] {
				consistent = false
				break
			}
		}
		if consistent {
			count++
		}
	}

	return count
}
