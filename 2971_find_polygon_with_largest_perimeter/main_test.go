package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestPerimeter(t *testing.T) {
	t.Run("should return 15", func(t *testing.T) {
		nums := []int{5, 5, 5}
		expected := int64(15)

		assert.Equal(t, expected, largestPerimeter(nums))
	})

	t.Run("should return 12", func(t *testing.T) {
		nums := []int{1, 12, 1, 2, 5, 50, 3}
		expected := int64(12)

		assert.Equal(t, expected, largestPerimeter(nums))
	})

	t.Run("should return -1", func(t *testing.T) {
		nums := []int{5, 5, 50}
		expected := int64(-1)

		assert.Equal(t, expected, largestPerimeter(nums))
	})

}
