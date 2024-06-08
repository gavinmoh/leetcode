package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckSubarraySum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{23, 2, 4, 6, 7}
		k := 6

		assert.True(t, checkSubarraySum(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{23, 2, 6, 4, 7}
		k := 6

		assert.True(t, checkSubarraySum(nums, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{23, 2, 6, 4, 7}
		k := 13

		assert.False(t, checkSubarraySum(nums, k))
	})

	t.Run("test case 4", func(t *testing.T) {
		nums := []int{23, 2, 4, 6, 6}
		k := 7

		assert.True(t, checkSubarraySum(nums, k))
	})

	t.Run("test case 5", func(t *testing.T) {
		nums := []int{1, 0}
		k := 2

		assert.False(t, checkSubarraySum(nums, k))
	})

	t.Run("test case 6", func(t *testing.T) {
		nums := []int{5, 0, 0, 0}
		k := 3

		assert.True(t, checkSubarraySum(nums, k))
	})
}
