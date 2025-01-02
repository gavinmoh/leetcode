package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVowelStrings(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		words := []string{"aba", "bcb", "ece", "aa", "e"}
		queries := [][]int{{0, 2}, {1, 4}, {1, 1}}
		expected := []int{2, 3, 0}

		assert.Equal(t, expected, vowelStrings(words, queries))
	})

	t.Run("test case 2", func(t *testing.T) {
		words := []string{"a", "e", "i"}
		queries := [][]int{{0, 2}, {0, 1}, {2, 2}}
		expected := []int{3, 2, 1}

		assert.Equal(t, expected, vowelStrings(words, queries))
	})
}
