package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxScoreWords(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		words := []string{"dog", "cat", "dad", "good"}
		letters := []byte("aacdddgoo")
		score := []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		expected := 23

		assert.Equal(t, expected, maxScoreWords(words, letters, score))
	})

	t.Run("test case 2", func(t *testing.T) {
		words := []string{"xxxz", "ax", "bx", "cx"}
		letters := []byte("zabcxxx")
		score := []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10}
		expected := 27

		assert.Equal(t, expected, maxScoreWords(words, letters, score))
	})

	t.Run("test case 3", func(t *testing.T) {
		words := []string{"leetcode"}
		letters := []byte("letcod")
		score := []int{0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}
		expected := 0

		assert.Equal(t, expected, maxScoreWords(words, letters, score))
	})
}
