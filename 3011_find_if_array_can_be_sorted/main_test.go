package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanSortArray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{8, 4, 2, 30, 15}
		assert.True(t, canSortArray(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		assert.True(t, canSortArray(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{3, 16, 8, 4, 2}
		assert.False(t, canSortArray(nums))
	})
}
