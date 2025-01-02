package main

func vowelStrings(words []string, queries [][]int) []int {
	prefixSum := make([]int, len(words)+1)
	currentSum := 0
	for i, word := range words {
		if isVowelString(word) {
			currentSum++
		}
		prefixSum[i+1] = currentSum
	}

	result := make([]int, len(queries))
	for i, query := range queries {
		left, right := query[0], query[1]
		result[i] = prefixSum[right+1] - prefixSum[left]
	}

	return result
}

func isVowelString(str string) bool {
	return isVowelChar(str[0]) && isVowelChar(str[len(str)-1])
}

func isVowelChar(char byte) bool {
	switch char {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}

	return false
}
