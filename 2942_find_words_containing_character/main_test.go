package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWordsContaining(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		words := []string{"leet", "code"}
		x := byte('e')
		expected := []int{0, 1}

		assert.Equal(t, expected, findWordsContaining(words, x))
	})

	t.Run("test case 2", func(t *testing.T) {
		words := []string{"abc", "bcd", "aaaa", "cbc"}
		x := byte('a')
		expected := []int{0, 2}

		assert.Equal(t, expected, findWordsContaining(words, x))
	})

	t.Run("test case 3", func(t *testing.T) {
		words := []string{"abc", "bcd", "aaaa", "cbc"}
		x := byte('z')
		expected := []int{}

		assert.Equal(t, expected, findWordsContaining(words, x))
	})
}
