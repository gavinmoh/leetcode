package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeightShelves(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		books := [][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}
		shelfWidth := 4
		expected := 6

		assert.Equal(t, expected, minHeightShelves(books, shelfWidth))
	})

	t.Run("test case 2", func(t *testing.T) {
		books := [][]int{{1, 3}, {2, 4}, {3, 2}}
		shelfWidth := 6
		expected := 4

		assert.Equal(t, expected, minHeightShelves(books, shelfWidth))
	})
}
