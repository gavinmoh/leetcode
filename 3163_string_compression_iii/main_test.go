package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompressedString(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		word := "abcde"
		expected := "1a1b1c1d1e"

		assert.Equal(t, expected, compressedString(word))
	})

	t.Run("test case 2", func(t *testing.T) {
		word := "aaaaaaaaaaaaaabb"
		expected := "9a5a2b"

		assert.Equal(t, expected, compressedString(word))
	})
}
