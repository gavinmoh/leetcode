package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWordsInLongestSubsequence(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		words := []string{"bab", "dab", "cab"}
		groups := []int{1, 2, 2}
		expected := []string{"bab", "dab"}

		assert.Equal(t, expected, getWordsInLongestSubsequence(words, groups))
	})

	t.Run("test case 2", func(t *testing.T) {
		words := []string{"a", "b", "c", "d"}
		groups := []int{1, 2, 3, 4}
		expected := []string{"a", "b", "c", "d"}

		assert.Equal(t, expected, getWordsInLongestSubsequence(words, groups))
	})
}
