package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShiftingLetters(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "abc"
		shifts := [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}}
		expected := "ace"

		assert.Equal(t, expected, shiftingLetters(s, shifts))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "dztz"
		shifts := [][]int{{0, 0, 0}, {1, 1, 1}}
		expected := "catz"

		assert.Equal(t, expected, shiftingLetters(s, shifts))
	})
}
