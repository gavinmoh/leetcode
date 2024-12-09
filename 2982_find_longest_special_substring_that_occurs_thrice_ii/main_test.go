package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumLength(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "aaaa"
		expected := 2

		assert.Equal(t, expected, maximumLength(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "abcdef"
		expected := -1

		assert.Equal(t, expected, maximumLength(s))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "abcaba"
		expected := 1

		assert.Equal(t, expected, maximumLength(s))
	})
}
