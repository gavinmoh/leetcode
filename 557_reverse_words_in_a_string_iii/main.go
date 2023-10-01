package main

func reverseWords(s string) string {
	strMap := make(map[int]string)

	for i, v := range s {
		strMap[i] = string(v)
	}

	wordMap := make(map[int]string)
	wordIndex := 0
	for i := 0; i < len(strMap); i++ {
		if strMap[i] == " " {
			wordIndex++
			continue
		}

		if wordMap[wordIndex] == "" {
			wordMap[wordIndex] = strMap[i]
		} else {
			wordMap[wordIndex] = strMap[i] + wordMap[wordIndex]
		}
	}

	result := ""
	for i := 0; i < len(wordMap); i++ {
		result += wordMap[i] + " "
	}

	return result[:len(result)-1]

}
