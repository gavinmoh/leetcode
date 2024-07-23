package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrequencySort(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 1, 2, 2, 2, 3}
		expected := []int{3, 1, 1, 2, 2, 2}

		assert.Equal(t, expected, frequencySort(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 3, 1, 3, 2}
		expected := []int{1, 3, 3, 2, 2}

		assert.Equal(t, expected, frequencySort(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}
		expected := []int{5, -1, 4, 4, -6, -6, 1, 1, 1}

		assert.Equal(t, expected, frequencySort(nums))
	})
}
