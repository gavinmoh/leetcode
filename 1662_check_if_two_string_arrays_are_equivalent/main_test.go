package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayStringsAreEqual(t *testing.T) {
	t.Run("should return true", func(t *testing.T) {
		word1 := []string{"ab", "c"}
		word2 := []string{"a", "bc"}
		expected := true
		assert.Equal(t, expected, arrayStringsAreEqual(word1, word2))
	})

	t.Run("should return false", func(t *testing.T) {
		word1 := []string{"a", "cb"}
		word2 := []string{"a", "bc"}
		expected := false
		assert.Equal(t, expected, arrayStringsAreEqual(word1, word2))
	})

	t.Run("should return true", func(t *testing.T) {
		word1 := []string{"abc", "d", "defg"}
		word2 := []string{"abcddefg"}
		expected := true
		assert.Equal(t, expected, arrayStringsAreEqual(word1, word2))
	})
}
