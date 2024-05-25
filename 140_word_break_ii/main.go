package main

import (
	"math"
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	n := len(s)
	wordMap := make(map[string]bool)
	minLength, maxLength := math.MaxInt, math.MinInt
	for _, word := range wordDict {
		wordMap[word] = true
		wordLength := len(word)
		if wordLength > maxLength {
			maxLength = wordLength
		}
		if wordLength < minLength {
			minLength = wordLength
		}
	}

	var dfs func(left int) [][]string
	dfs = func(left int) [][]string {
		if left >= n {
			return [][]string{{}}
		}

		result := [][]string{}
		for right := left + 1; right <= n; right++ {
			substr := s[left:right]
			if len(substr) < minLength {
				continue
			}

			if len(substr) > maxLength {
				break
			}

			if !wordMap[substr] {
				continue
			}

			// add space
			sentences := dfs(right)
			for _, sentence := range sentences {
				newSentence := append([]string{substr}, sentence...)
				result = append(result, newSentence)
			}
		}

		return result
	}

	outputs := []string{}
	for _, sentence := range dfs(0) {
		outputs = append(outputs, strings.Join(sentence, " "))
	}

	return outputs
}
