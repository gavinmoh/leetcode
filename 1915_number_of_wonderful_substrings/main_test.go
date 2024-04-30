package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWonderfulSubstrings(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		word := "aba"
		expected := int64(4)

		assert.Equal(t, expected, wonderfulSubstrings(word))
	})

	t.Run("test case 2", func(t *testing.T) {
		word := "aabb"
		expected := int64(9)

		assert.Equal(t, expected, wonderfulSubstrings(word))
	})

	t.Run("test case 3", func(t *testing.T) {
		word := "he"
		expected := int64(2)

		assert.Equal(t, expected, wonderfulSubstrings(word))
	})
}
