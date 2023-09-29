package main

func minExtraChar(s string, dictionary []string) int {
	// create a lookup map for dictionary
	dictionaryMap := make(map[string]bool)
	for _, word := range dictionary {
		dictionaryMap[word] = true
	}

	n := len(s)
	minExtraCount := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minExtraCount[i] = minExtraCount[i-1] + 1
		for j := 0; j < i; j++ {
			lookupWord := s[j:i]
			if dictionaryMap[lookupWord] && minExtraCount[j] < minExtraCount[i] {
				minExtraCount[i] = minExtraCount[j]
			}
		}
	}
	return minExtraCount[n]
}
