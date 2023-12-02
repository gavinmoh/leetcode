package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCharacters(t *testing.T) {
	t.Run("should return 6", func(t *testing.T) {
		words := []string{"cat", "bt", "hat", "tree"}
		chars := "atach"
		expected := 6
		assert.Equal(t, expected, countCharacters(words, chars))
	})

	t.Run("should return 10", func(t *testing.T) {
		words := []string{"hello", "world", "leetcode"}
		chars := "welldonehoneyr"
		expected := 10
		assert.Equal(t, expected, countCharacters(words, chars))
	})
}
