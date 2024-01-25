package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubsequence(t *testing.T) {
	t.Run("should return 3", func(t *testing.T) {
		text1 := "abcde"
		text2 := "ace"
		expected := 3
		assert.Equal(t, expected, longestCommonSubsequence(text1, text2))
	})

	t.Run("should return 3", func(t *testing.T) {
		text1 := "abc"
		text2 := "abc"
		expected := 3
		assert.Equal(t, expected, longestCommonSubsequence(text1, text2))
	})

	t.Run("should return 0", func(t *testing.T) {
		text1 := "abc"
		text2 := "def"
		expected := 0
		assert.Equal(t, expected, longestCommonSubsequence(text1, text2))
	})

	t.Run("should return 1", func(t *testing.T) {
		text1 := "bsbininm"
		text2 := "jmjkbkjkv"
		expected := 1
		assert.Equal(t, expected, longestCommonSubsequence(text1, text2))
	})
}
