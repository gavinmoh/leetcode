package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSubrraysWithSum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 0, 1, 0, 1}
		goal := 2
		expected := 4

		assert.Equal(t, expected, numSubarraysWithSum(nums, goal))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{0, 0, 0, 0, 0}
		goal := 0
		expected := 15

		assert.Equal(t, expected, numSubarraysWithSum(nums, goal))
	})
}
