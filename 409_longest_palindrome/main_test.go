package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "abccccdd"
		expected := 7

		assert.Equal(t, expected, longestPalindrome(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "a"
		expected := 1

		assert.Equal(t, expected, longestPalindrome(s))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "bananas"
		expected := 5

		assert.Equal(t, expected, longestPalindrome(s))
	})

	t.Run("test case 4", func(t *testing.T) {
		s := "bb"
		expected := 2

		assert.Equal(t, expected, longestPalindrome(s))
	})
}
