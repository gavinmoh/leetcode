package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPivotArray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{9, 12, 5, 10, 14, 3, 10}
		pivot := 10
		expected := []int{9, 5, 3, 10, 10, 12, 14}

		assert.Equal(t, expected, pivotArray(nums, pivot))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{-3, 4, 3, 2}
		pivot := 2
		expected := []int{-3, 2, 4, 3}

		assert.Equal(t, expected, pivotArray(nums, pivot))
	})
}
