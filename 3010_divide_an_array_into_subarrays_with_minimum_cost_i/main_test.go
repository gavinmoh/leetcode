package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumCost(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 12}
		expected := 6

		assert.Equal(t, expected, minimumCost(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{5, 4, 3}
		expected := 12

		assert.Equal(t, expected, minimumCost(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{10, 3, 1, 1}
		expected := 12

		assert.Equal(t, expected, minimumCost(nums))
	})
}
