package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountOfSubstrings(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		word := "aeioqq"
		k := 1
		expected := int64(0)

		assert.Equal(t, expected, countOfSubstrings(word, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		word := "aeiou"
		k := 0
		expected := int64(1)

		assert.Equal(t, expected, countOfSubstrings(word, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		word := "ieaouqqieaouqq"
		k := 1
		expected := int64(3)

		assert.Equal(t, expected, countOfSubstrings(word, k))
	})
}
