package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestBalanced(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "abbac"
		expected := 4

		assert.Equal(t, expected, longestBalanced(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "zzabccy"
		expected := 4

		assert.Equal(t, expected, longestBalanced(s))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "aba"
		expected := 2

		assert.Equal(t, expected, longestBalanced(s))
	})
}
