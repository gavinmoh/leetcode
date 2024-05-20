package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsetXORSum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 3}
		expected := 6

		assert.Equal(t, expected, subsetXORSum(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{5, 1, 6}
		expected := 28

		assert.Equal(t, expected, subsetXORSum(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{3, 4, 5, 6, 7, 8}
		expected := 480

		assert.Equal(t, expected, subsetXORSum(nums))
	})
}
