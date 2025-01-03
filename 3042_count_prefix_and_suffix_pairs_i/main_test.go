package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPrefixSuffixPairs(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		words := []string{"a", "aba", "ababa", "aa"}
		expected := 4

		assert.Equal(t, expected, countPrefixSuffixPairs(words))
	})

	t.Run("test case 2", func(t *testing.T) {
		words := []string{"pa", "papa", "ma", "mama"}
		expected := 2

		assert.Equal(t, expected, countPrefixSuffixPairs(words))
	})

	t.Run("test case 3", func(t *testing.T) {
		words := []string{"abab", "ab"}
		expected := 0

		assert.Equal(t, expected, countPrefixSuffixPairs(words))
	})
}
