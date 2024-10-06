package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreSentencesSimilar(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		sentence1 := "My name is Haley"
		sentence2 := "My Haley"

		assert.True(t, areSentencesSimilar(sentence1, sentence2))
	})

	t.Run("test case 2", func(t *testing.T) {
		sentence1 := "of"
		sentence2 := "A lot of words"

		assert.False(t, areSentencesSimilar(sentence1, sentence2))
	})

	t.Run("test case 3", func(t *testing.T) {
		sentence1 := "Eating right now"
		sentence2 := "Eating"

		assert.True(t, areSentencesSimilar(sentence1, sentence2))
	})

	t.Run("test case 4", func(t *testing.T) {
		sentence1 := "qbaVXO Msgr aEWD v ekcb"
		sentence2 := "Msgr aEWD ekcb"

		assert.False(t, areSentencesSimilar(sentence1, sentence2))
	})
}
