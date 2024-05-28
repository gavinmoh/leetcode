package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualSubstring(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s1 := "abcd"
		s2 := "bcdf"
		maxCost := 3
		expected := 3

		assert.Equal(t, expected, equalSubstring(s1, s2, maxCost))
	})

	t.Run("test case 2", func(t *testing.T) {
		s1 := "abcd"
		s2 := "cdef"
		maxCost := 3
		expected := 1

		assert.Equal(t, expected, equalSubstring(s1, s2, maxCost))
	})

	t.Run("test case 3", func(t *testing.T) {
		s1 := "abcd"
		s2 := "acde"
		maxCost := 0
		expected := 1

		assert.Equal(t, expected, equalSubstring(s1, s2, maxCost))
	})

	t.Run("test case 4", func(t *testing.T) {
		s1 := "abcd"
		s2 := "abcd"
		maxCost := 0
		expected := 4

		assert.Equal(t, expected, equalSubstring(s1, s2, maxCost))
	})
}
