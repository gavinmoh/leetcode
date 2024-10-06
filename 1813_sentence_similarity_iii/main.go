package main

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	words1 := stringToSlice(sentence1)
	words2 := stringToSlice(sentence2)

	for len(words1) > 0 && len(words2) > 0 {
		if words1[0] == words2[0] {
			words1 = words1[1:]
			words2 = words2[1:]
		} else {
			break
		}
	}

	for len(words1) > 0 && len(words2) > 0 {
		if words1[len(words1)-1] == words2[len(words2)-1] {
			words1 = words1[:len(words1)-1]
			words2 = words2[:len(words2)-1]
		} else {
			break
		}
	}

	if len(words1) == 0 || len(words2) == 0 {
		return true
	}

	return false
}

func stringToSlice(s string) []string {
	result := []string{}
	word := []rune{}
	for _, letter := range s {
		if letter == ' ' {
			result = append(result, string(word))
			word = []rune{}
		} else {
			word = append(word, letter)
		}
	}
	result = append(result, string(word))

	return result
}
