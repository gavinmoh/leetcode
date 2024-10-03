package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSubarray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{3, 1, 4, 2}
		p := 6
		expected := 1

		assert.Equal(t, expected, minSubarray(nums, p))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{6, 3, 5, 2}
		p := 9
		expected := 2

		assert.Equal(t, expected, minSubarray(nums, p))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{1, 2, 3}
		p := 3
		expected := 0

		assert.Equal(t, expected, minSubarray(nums, p))
	})
}
