package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResultsArray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 3, 2, 5}
		k := 3
		expected := []int{3, 4, -1, -1, -1}

		assert.Equal(t, expected, resultsArray(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 2, 2, 2, 2}
		k := 4
		expected := []int{-1, -1}

		assert.Equal(t, expected, resultsArray(nums, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{3, 2, 3, 2, 3, 2}
		k := 2
		expected := []int{-1, 3, -1, 3, -1}

		assert.Equal(t, expected, resultsArray(nums, k))
	})
}
