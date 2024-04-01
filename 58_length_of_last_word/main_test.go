package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "Hello World"
		expected := 5

		assert.Equal(t, expected, lengthOfLastWord(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "   fly me   to   the moon  "
		expected := 4

		assert.Equal(t, expected, lengthOfLastWord(s))
	})

	t.Run("test case 1", func(t *testing.T) {
		s := "luffy is still joyboy"
		expected := 6

		assert.Equal(t, expected, lengthOfLastWord(s))
	})
}
