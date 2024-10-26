package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsCircularSentence(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		sentence := "leetcode exercises sound delightful"

		assert.True(t, isCircularSentence(sentence))
	})

	t.Run("test case 2", func(t *testing.T) {
		sentence := "eetcode"

		assert.True(t, isCircularSentence(sentence))
	})

	t.Run("test case 3", func(t *testing.T) {
		sentence := "Leetcode is cool"

		assert.False(t, isCircularSentence(sentence))
	})
}
