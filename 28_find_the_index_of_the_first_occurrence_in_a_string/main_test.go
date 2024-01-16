package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	t.Run("should return 0", func(t *testing.T) {
		haystack := "sadbutsad"
		needle := "sad"
		expected := 0
		assert.Equal(t, expected, strStr(haystack, needle))
	})

	t.Run("should return -1", func(t *testing.T) {
		haystack := "leetcode"
		needle := "leeto"
		expected := -1
		assert.Equal(t, expected, strStr(haystack, needle))
	})

	t.Run("should return 0", func(t *testing.T) {
		haystack := "a"
		needle := "a"
		expected := 0
		assert.Equal(t, expected, strStr(haystack, needle))
	})
}
