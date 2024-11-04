package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestCombination(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		candidates := []int{16, 17, 71, 62, 12, 24, 14}
		expected := 4

		assert.Equal(t, expected, largestCombination(candidates))
	})

	t.Run("test case 2", func(t *testing.T) {
		candidates := []int{8, 8}
		expected := 2

		assert.Equal(t, expected, largestCombination(candidates))
	})
}
