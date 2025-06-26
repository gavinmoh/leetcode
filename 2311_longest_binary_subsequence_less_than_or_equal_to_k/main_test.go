package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubsequence(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "1001010"
		k := 5
		expected := 5

		assert.Equal(t, expected, longestSubsequence(s, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "00101001"
		k := 1
		expected := 6

		assert.Equal(t, expected, longestSubsequence(s, k))
	})
}
