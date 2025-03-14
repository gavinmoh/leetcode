package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumCandies(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		candies := []int{5, 8, 6}
		k := int64(3)
		expected := 5

		assert.Equal(t, expected, maximumCandies(candies, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		candies := []int{2, 5}
		k := int64(11)
		expected := 0

		assert.Equal(t, expected, maximumCandies(candies, k))
	})
}
