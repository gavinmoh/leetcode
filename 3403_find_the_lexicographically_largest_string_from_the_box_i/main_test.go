package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnswerString(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		word := "dbca"
		numFriends := 2
		expected := "dbc"

		assert.Equal(t, expected, answerString(word, numFriends))
	})

	t.Run("test case 2", func(t *testing.T) {
		word := "gggg"
		numFriends := 4
		expected := "g"

		assert.Equal(t, expected, answerString(word, numFriends))
	})

	t.Run("test case 3", func(t *testing.T) {
		word := "aann"
		numFriends := 2
		expected := "nn"

		assert.Equal(t, expected, answerString(word, numFriends))
	})

	t.Run("test case 4", func(t *testing.T) {
		word := "gh"
		numFriends := 1
		expected := "gh"

		assert.Equal(t, expected, answerString(word, numFriends))
	})
}
