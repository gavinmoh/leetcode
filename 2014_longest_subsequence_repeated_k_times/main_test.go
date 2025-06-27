package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubsequenceRepeatedK(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "letsleetcode"
		k := 2
		expected := "let"

		assert.Equal(t, expected, longestSubsequenceRepeatedK(s, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "bb"
		k := 2
		expected := "b"

		assert.Equal(t, expected, longestSubsequenceRepeatedK(s, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "ab"
		k := 2
		expected := ""

		assert.Equal(t, expected, longestSubsequenceRepeatedK(s, k))
	})
}
