package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFinalPrices(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		prices := []int{8, 4, 6, 2, 3}
		expected := []int{4, 2, 4, 2, 3}

		assert.Equal(t, expected, finalPrices(prices))
	})

	t.Run("test case 2", func(t *testing.T) {
		prices := []int{1, 2, 3, 4, 5}
		expected := []int{1, 2, 3, 4, 5}

		assert.Equal(t, expected, finalPrices(prices))
	})

	t.Run("test case 3", func(t *testing.T) {
		prices := []int{10, 1, 1, 6}
		expected := []int{9, 0, 1, 6}

		assert.Equal(t, expected, finalPrices(prices))
	})
}
