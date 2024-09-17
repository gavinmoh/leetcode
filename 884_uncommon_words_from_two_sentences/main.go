package main

func uncommonFromSentences(s1 string, s2 string) []string {
	words := make(map[string]int)
	for _, word := range sentenceToWords(s1) {
		words[word] += 1
	}
	for _, word := range sentenceToWords(s2) {
		words[word] += 1
	}

	results := []string{}
	for word, count := range words {
		if count == 1 {
			results = append(results, word)
		}
	}

	return results
}

func sentenceToWords(sentence string) []string {
	word := ""
	words := []string{}
	for _, letter := range sentence {
		if letter == ' ' {
			words = append(words, word)
			word = ""
		} else {
			word += string(letter)
		}
	}
	words = append(words, word)

	return words
}
