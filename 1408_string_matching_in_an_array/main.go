package main

import "strings"

func stringMatching(words []string) []string {
	resultMap := map[string]bool{}

	for _, substr := range words {
		for _, word := range words {
			if word == substr {
				continue
			}

			if _, ok := resultMap[substr]; ok {
				continue
			}

			if strings.Contains(word, substr) {
				resultMap[substr] = true
			}
		}
	}

	result := []string{}
	for substr, _ := range resultMap {
		result = append(result, substr)
	}

	return result
}
