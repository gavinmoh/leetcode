package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxCoins(t *testing.T) {
	t.Run("should return 9", func(t *testing.T) {
		piles := []int{2, 4, 1, 2, 7, 8}
		expected := 9
		assert.Equal(t, expected, maxCoins(piles))
	})

	t.Run("should return 4", func(t *testing.T) {
		piles := []int{2, 4, 5}
		expected := 4
		assert.Equal(t, expected, maxCoins(piles))
	})

	t.Run("should return 18", func(t *testing.T) {
		piles := []int{9, 8, 7, 6, 5, 1, 2, 3, 4}
		expected := 18
		assert.Equal(t, expected, maxCoins(piles))
	})

}
