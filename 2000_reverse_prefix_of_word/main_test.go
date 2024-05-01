package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReversePrefix(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		word := "abcdefd"
		ch := byte('d')
		expected := "dcbaefd"

		assert.Equal(t, expected, reversePrefix(word, ch))
	})

	t.Run("test case 2", func(t *testing.T) {
		word := "xyxzxe"
		ch := byte('z')
		expected := "zxyxxe"

		assert.Equal(t, expected, reversePrefix(word, ch))
	})

	t.Run("test case 3", func(t *testing.T) {
		word := "abcd"
		ch := byte('z')
		expected := "abcd"

		assert.Equal(t, expected, reversePrefix(word, ch))
	})
}
