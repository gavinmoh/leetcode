package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxFrequencyElements(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 2, 3, 1, 4}
		expected := 4

		assert.Equal(t, expected, maxFrequencyElements(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		expected := 5

		assert.Equal(t, expected, maxFrequencyElements(nums))
	})
}
