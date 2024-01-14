package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloseStrings(t *testing.T) {
	t.Run("should return true", func(t *testing.T) {
		word1 := "abc"
		word2 := "bca"
		assert.True(t, closeStrings(word1, word2))
	})

	t.Run("should return false", func(t *testing.T) {
		word1 := "a"
		word2 := "aa"
		assert.False(t, closeStrings(word1, word2))
	})

	t.Run("should return true", func(t *testing.T) {
		word1 := "cabbba"
		word2 := "abbccc"
		assert.True(t, closeStrings(word1, word2))
	})
}
