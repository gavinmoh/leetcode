package main

func numWays(words []string, target string) int {
	dp := make([][]int, len(words[0]))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(target))
		for j := 0; j < len(target); j++ {
			dp[i][j] = -1
		}
	}

	charFreq := make([][]int, len(words[0]))
	for i := 0; i < len(charFreq); i++ {
		charFreq[i] = make([]int, 26)
	}

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[0]); j++ {
			char := words[i][j] - 'a'
			charFreq[j][char]++
		}
	}

	var getWords func(wordsIdx, targetIdx int) int
	getWords = func(wordsIdx, targetIdx int) int {
		if targetIdx == len(target) {
			return 1
		}

		if wordsIdx == len(words[0]) || (len(words[0])-wordsIdx) < (len(target)-targetIdx) {
			return 0
		}

		if dp[wordsIdx][targetIdx] != -1 {
			return dp[wordsIdx][targetIdx]
		}

		count := 0
		current := target[targetIdx] - 'a'
		count += getWords(wordsIdx+1, targetIdx)
		count += charFreq[wordsIdx][current] * getWords(wordsIdx+1, targetIdx+1)

		dp[wordsIdx][targetIdx] = count % 1000000007
		return dp[wordsIdx][targetIdx]
	}

	return getWords(0, 0)
}
