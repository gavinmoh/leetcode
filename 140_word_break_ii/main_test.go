package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordBreak(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "catsanddog"
		wordDict := []string{"cat", "cats", "and", "sand", "dog"}
		expectedOutputs := []string{"cats and dog", "cat sand dog"}
		actualOutputs := wordBreak(s, wordDict)

		assert.Equal(t, len(expectedOutputs), len(actualOutputs))
		for _, output := range expectedOutputs {
			assert.Contains(t, actualOutputs, output)
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "pineapplepenapple"
		wordDict := []string{"apple", "pen", "applepen", "pine", "pineapple"}
		expectedOutputs := []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"}
		actualOutputs := wordBreak(s, wordDict)

		assert.Equal(t, len(expectedOutputs), len(actualOutputs))
		for _, output := range expectedOutputs {
			assert.Contains(t, actualOutputs, output)
		}
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "catsandog"
		wordDict := []string{"cats", "dog", "sand", "and", "cat"}
		expectedOutputs := []string{}
		actualOutputs := wordBreak(s, wordDict)

		assert.Equal(t, len(expectedOutputs), len(actualOutputs))
		for _, output := range expectedOutputs {
			assert.Contains(t, actualOutputs, output)
		}
	})
}
