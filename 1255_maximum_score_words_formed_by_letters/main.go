package main

func maxScoreWords(words []string, letters []byte, score []int) int {
	n := len(words)
	lettersCount := make(map[byte]int)
	for _, letter := range letters {
		lettersCount[letter] += 1
	}

	scores := make(map[string]int)
	for _, word := range words {
		currentScore := 0
		for _, char := range word {
			currentScore += score[char-'a']
		}
		scores[word] = currentScore
	}

	var dfs func(left int) int
	dfs = func(left int) int {
		if left >= n {
			return 0
		}

		// skip
		result := dfs(left + 1)

		// include
		word := words[left]
		enoughLetters := true
		for _, char := range word {
			lettersCount[byte(char)] -= 1
			if lettersCount[byte(char)] < 0 {
				enoughLetters = false
			}
		}

		if enoughLetters {
			result = max(result, scores[word]+dfs(left+1))
		}

		// uncount the letters
		for _, char := range word {
			lettersCount[byte(char)] += 1
		}

		return result
	}

	return dfs(0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
