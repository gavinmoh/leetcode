package main

func commonChars(words []string) []string {
	commonStr := []rune{}
	for _, letter := range words[0] {
		commonStr = append(commonStr, letter)
	}

	for i := 1; i < len(words); i++ {
		commonStr = intersect(commonStr, words[i])
	}

	result := []string{}
	for _, letter := range commonStr {
		result = append(result, string(letter))
	}

	return result
}

func intersect(s1 []rune, s2 string) []rune {
	output := []rune{}
	letterCount := make(map[rune]int)
	for _, char := range s1 {
		letterCount[char]++
	}

	for _, char := range s2 {
		if count, ok := letterCount[char]; ok && count > 0 {
			output = append(output, char)
			letterCount[char]--
		}
	}

	return output
}
