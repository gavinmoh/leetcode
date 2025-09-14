package main

func spellchecker(wordlist []string, queries []string) []string {
	result := make([]string, len(queries))

	wordMap := make(map[string]bool)
	for _, word := range wordlist {
		wordMap[word] = true
	}

	for i, query := range queries {
		if _, ok := wordMap[query]; ok {
			result[i] = query
			continue
		}

		vowelMatch := ""
		for _, word := range wordlist {
			if len(word) != len(query) {
				continue
			}

			if isCapitalizationMatch(word, query) {
				result[i] = word
				break
			}

			if isVowelMatch(word, query) && vowelMatch == "" {
				vowelMatch = word
			}
		}

		if result[i] != "" {
			continue
		}

		result[i] = vowelMatch
	}

	return result
}

func isVowelMatch(word, query string) bool {
	for i := 0; i < len(word); i++ {
		c1, c2 := rune(word[i]), rune(query[i])
		if c1 == c2 {
			continue
		}

		if isVowel(c1) && isVowel(c2) {
			continue
		}

		// check for capitalization
		if abs(int(c1-c2)) == 32 {
			continue
		}

		return false
	}

	return true
}

func isCapitalizationMatch(word, query string) bool {
	for i := 0; i < len(word); i++ {
		c1, c2 := rune(word[i]), rune(query[i])
		if c1 == c2 {
			continue
		}

		// check for capitalization
		if abs(int(c1-c2)) == 32 {
			continue
		}

		return false
	}

	return true
}

func isVowel(c rune) bool {
	switch c {
	case 'A', 'E', 'I', 'O', 'U':
		fallthrough
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
