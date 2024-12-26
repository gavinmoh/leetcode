package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTargetSumWays(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 1, 1, 1, 1}
		target := 3
		expected := 5

		assert.Equal(t, expected, findTargetSumWays(nums, target))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1}
		target := 1
		expected := 1

		assert.Equal(t, expected, findTargetSumWays(nums, target))
	})
}
