package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoneGameII(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		piles := []int{2, 7, 9, 4, 4}
		expected := 10

		assert.Equal(t, expected, stoneGameII(piles))
	})

	t.Run("test case 2", func(t *testing.T) {
		piles := []int{1, 2, 3, 4, 5, 100}
		expected := 104

		assert.Equal(t, expected, stoneGameII(piles))
	})
}
