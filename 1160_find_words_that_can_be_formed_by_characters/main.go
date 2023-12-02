package main

func countCharacters(words []string, chars string) int {
	charMap := make(map[rune]int, len(chars))
	output := 0

	for _, char := range chars {
		charMap[char] += 1
	}

	for _, word := range words {
		wordMap := make(map[rune]int, len(word))

		for _, char := range word {
			wordMap[char] += 1
		}

		formable := true
		for char, freq := range wordMap {
			count, ok := charMap[char]
			if !ok || freq > count {
				formable = false
				break
			}
		}

		if !formable {
			continue
		}

		output += len(word)
	}

	return output
}
