package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBagOfTokensScore(t *testing.T) {
	t.Run("should return 0", func(t *testing.T) {
		tokens := []int{100}
		power := 50
		expected := 0

		assert.Equal(t, expected, bagOfTokensScore(tokens, power))
	})

	t.Run("should return 1", func(t *testing.T) {
		tokens := []int{200, 100}
		power := 150
		expected := 1

		assert.Equal(t, expected, bagOfTokensScore(tokens, power))
	})

	t.Run("should return 2", func(t *testing.T) {
		tokens := []int{100, 200, 300, 400}
		power := 200
		expected := 2

		assert.Equal(t, expected, bagOfTokensScore(tokens, power))
	})

	t.Run("should return 3", func(t *testing.T) {
		tokens := []int{33, 4, 28, 24, 96}
		power := 35
		expected := 3

		assert.Equal(t, expected, bagOfTokensScore(tokens, power))
	})

	t.Run("should return 6", func(t *testing.T) {
		tokens := []int{59, 0, 15, 33, 79, 72, 34, 62, 4, 16}
		power := 99
		expected := 6

		assert.Equal(t, expected, bagOfTokensScore(tokens, power))
	})
}
