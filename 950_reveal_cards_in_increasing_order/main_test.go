package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeckRevealedIncreasing(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		deck := []int{17, 13, 11, 2, 3, 5, 7}
		expected := []int{2, 13, 3, 11, 5, 17, 7}

		assert.Equal(t, expected, deckRevealedIncreasing(deck))
	})

	t.Run("test case 2", func(t *testing.T) {
		deck := []int{1000, 1}
		expected := []int{1, 1000}

		assert.Equal(t, expected, deckRevealedIncreasing(deck))
	})
}
