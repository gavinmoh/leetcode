package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestDivisibleSubset(t *testing.T) {
	t.Run("should return [1,2]", func(t *testing.T) {
		nums := []int{1, 2, 3}
		expected := []int{1, 2}

		assert.Equal(t, expected, largestDivisibleSubset(nums))
	})

	t.Run("should return [1,2,4,8]", func(t *testing.T) {
		nums := []int{1, 2, 4, 8}
		expected := []int{1, 2, 4, 8}

		assert.Equal(t, expected, largestDivisibleSubset(nums))
	})
}
