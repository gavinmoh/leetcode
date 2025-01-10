package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordSubsets(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		words1 := []string{"amazon", "apple", "facebook", "google", "leetcode"}
		words2 := []string{"e", "o"}
		expected := []string{"facebook", "google", "leetcode"}

		assert.Equal(t, expected, wordSubsets(words1, words2))
	})

	t.Run("test case 2", func(t *testing.T) {
		words1 := []string{"amazon", "apple", "facebook", "google", "leetcode"}
		words2 := []string{"l", "e"}
		expected := []string{"apple", "google", "leetcode"}

		assert.Equal(t, expected, wordSubsets(words1, words2))
	})

	t.Run("test case 3", func(t *testing.T) {
		words1 := []string{"acaac", "cccbb", "aacbb", "caacc", "bcbbb"}
		words2 := []string{"c", "cc", "b"}
		expected := []string{"cccbb"}

		assert.Equal(t, expected, wordSubsets(words1, words2))
	})
}
