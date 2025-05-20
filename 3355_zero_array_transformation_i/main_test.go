package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsZeroArray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 0, 1}
		queries := [][]int{{0, 2}}

		assert.True(t, isZeroArray(nums, queries))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{4, 3, 2, 1}
		queries := [][]int{{1, 3}, {0, 2}}

		assert.False(t, isZeroArray(nums, queries))
	})
}
