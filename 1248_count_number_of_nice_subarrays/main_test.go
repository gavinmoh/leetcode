package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfSubarrays(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 1, 2, 1, 1}
		k := 3
		expected := 2

		assert.Equal(t, expected, numberOfSubarrays(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 4, 6}
		k := 1
		expected := 0

		assert.Equal(t, expected, numberOfSubarrays(nums, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}
		k := 2
		expected := 16

		assert.Equal(t, expected, numberOfSubarrays(nums, k))
	})
}
