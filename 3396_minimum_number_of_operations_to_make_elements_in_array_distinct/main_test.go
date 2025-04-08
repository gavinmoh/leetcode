package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumOperations(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 2, 3, 3, 5, 7}
		expected := 2

		assert.Equal(t, expected, minimumOperations(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{4, 5, 6, 4, 4}
		expected := 2

		assert.Equal(t, expected, minimumOperations(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{6, 7, 8, 9}
		expected := 0

		assert.Equal(t, expected, minimumOperations(nums))
	})
}
