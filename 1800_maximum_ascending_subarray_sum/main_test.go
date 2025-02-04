package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxAscendingSum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{10, 20, 30, 5, 10, 50}
		expected := 65

		assert.Equal(t, expected, maxAscendingSum(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{10, 20, 30, 40, 50}
		expected := 150

		assert.Equal(t, expected, maxAscendingSum(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{12, 17, 15, 13, 10, 11, 12}
		expected := 33

		assert.Equal(t, expected, maxAscendingSum(nums))
	})

	t.Run("test case 4", func(t *testing.T) {
		nums := []int{3, 6, 10, 1, 8, 9, 9, 8, 9}
		expected := 19

		assert.Equal(t, expected, maxAscendingSum(nums))
	})
}
