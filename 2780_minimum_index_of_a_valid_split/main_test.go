package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumIndex(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 2, 2}
		expected := 2

		assert.Equal(t, expected, minimumIndex(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 1, 3, 1, 1, 1, 7, 1, 2, 1}
		expected := 4

		assert.Equal(t, expected, minimumIndex(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{3, 3, 3, 3, 7, 2, 2}
		expected := -1

		assert.Equal(t, expected, minimumIndex(nums))
	})
}
