package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrefixOfWord(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		sentence := "i love eating burger"
		searchWord := "burg"
		expected := 4

		assert.Equal(t, expected, isPrefixOfWord(sentence, searchWord))
	})

	t.Run("test case 2", func(t *testing.T) {
		sentence := "this problem is an easy problem"
		searchWord := "pro"
		expected := 2

		assert.Equal(t, expected, isPrefixOfWord(sentence, searchWord))
	})

	t.Run("test case 3", func(t *testing.T) {
		sentence := "i am tired"
		searchWord := "you"
		expected := -1

		assert.Equal(t, expected, isPrefixOfWord(sentence, searchWord))
	})
}
