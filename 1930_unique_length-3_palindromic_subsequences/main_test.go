package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPalindromicSubsequence(t *testing.T) {
	t.Run("should return 3", func(t *testing.T) {
		s := "aabca"
		expected := 3
		assert.Equal(t, expected, countPalindromicSubsequence(s))
	})

	t.Run("should return 0", func(t *testing.T) {
		s := "adc"
		expected := 0
		assert.Equal(t, expected, countPalindromicSubsequence(s))
	})

	t.Run("should return 4", func(t *testing.T) {
		s := "bbcbaba"
		expected := 4
		assert.Equal(t, expected, countPalindromicSubsequence(s))
	})
}
