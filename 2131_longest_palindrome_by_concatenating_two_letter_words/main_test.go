package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		words := []string{"lc", "cl", "gg"}
		expected := 6

		assert.Equal(t, expected, longestPalindrome(words))
	})

	t.Run("test case 2", func(t *testing.T) {
		words := []string{"ab", "ty", "yt", "lc", "cl", "ab"}
		expected := 8

		assert.Equal(t, expected, longestPalindrome(words))
	})

	t.Run("test case 3", func(t *testing.T) {
		words := []string{"cc", "ll", "xx"}
		expected := 2

		assert.Equal(t, expected, longestPalindrome(words))
	})

	t.Run("test case 4", func(t *testing.T) {
		words := []string{"dd", "aa", "bb", "dd", "aa", "dd", "bb", "dd", "aa", "cc", "bb", "cc", "dd", "cc"}
		expected := 22

		assert.Equal(t, expected, longestPalindrome(words))
	})
}
