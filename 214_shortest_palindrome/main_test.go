package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestPalindrome(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "aacecaaa"
		expected := "aaacecaaa"

		assert.Equal(t, expected, shortestPalindrome(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "abcd"
		expected := "dcbabcd"

		assert.Equal(t, expected, shortestPalindrome(s))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := ""
		expected := ""

		assert.Equal(t, expected, shortestPalindrome(s))
	})

	t.Run("test case 4", func(t *testing.T) {
		s := "a"
		expected := "a"

		assert.Equal(t, expected, shortestPalindrome(s))
	})

	t.Run("test case 5", func(t *testing.T) {
		s := "aba"
		expected := "aba"

		assert.Equal(t, expected, shortestPalindrome(s))
	})
}
