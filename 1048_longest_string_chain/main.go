package main

import "sort"

func longestStrChain(words []string) int {
	// Sort words by length
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	chain := make(map[string]int)
	count := 0
	for _, word := range words {
		chain[word] = 1
		for i := 0; i < len(word); i++ {
			// Remove one character from word
			// and check if it exists in chain
			// and if it's chain length is greater
			// than current chain length
			// if so, update chain length
			// for current word
			str := word[:i] + word[i+1:]
			if _, ok := chain[str]; ok {
				if chain[str]+1 > chain[word] {
					chain[word] = chain[str] + 1
				}
			}
		}

		if chain[word] > count {
			count = chain[word]
		}

	}

	return count
}
