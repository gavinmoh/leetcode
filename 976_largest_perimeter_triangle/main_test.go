package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestPerimeter(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{2, 1, 2}
		expected := 5

		assert.Equal(t, expected, largestPerimeter(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2, 1, 10}
		expected := 0

		assert.Equal(t, expected, largestPerimeter(nums))
	})
}
