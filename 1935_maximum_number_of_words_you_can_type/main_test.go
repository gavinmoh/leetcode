package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanBeTypedWords(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		text := "hello world"
		brokenLetters := "ad"
		expected := 1

		assert.Equal(t, expected, canBeTypedWords(text, brokenLetters))
	})

	t.Run("test case 2", func(t *testing.T) {
		text := "leet code"
		brokenLetters := "lt"
		expected := 1

		assert.Equal(t, expected, canBeTypedWords(text, brokenLetters))
	})

	t.Run("test case 3", func(t *testing.T) {
		text := "lett code"
		brokenLetters := "e"
		expected := 0

		assert.Equal(t, expected, canBeTypedWords(text, brokenLetters))
	})
}
