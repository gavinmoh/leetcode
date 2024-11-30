package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumSize(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{9}
		maxOperations := 2
		expected := 3

		assert.Equal(t, expected, minimumSize(nums, maxOperations))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 4, 8, 2}
		maxOperations := 4
		expected := 2

		assert.Equal(t, expected, minimumSize(nums, maxOperations))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{431, 922, 158, 60, 192, 14, 788, 146, 788, 775, 772, 792, 68, 143, 376, 375, 877, 516, 595, 82, 56, 704, 160, 403, 713, 504, 67, 332, 26}
		maxOperations := 80
		expected := 129

		assert.Equal(t, expected, minimumSize(nums, maxOperations))
	})
}
