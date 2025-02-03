package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumDeletions(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		word := "aabcaba"
		k := 0
		expected := 3

		assert.Equal(t, expected, minimumDeletions(word, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		word := "dabdcbdcdcd"
		k := 2
		expected := 2

		assert.Equal(t, expected, minimumDeletions(word, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		word := "aaabaaa"
		k := 2
		expected := 1

		assert.Equal(t, expected, minimumDeletions(word, k))
	})
}
