package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 1, 3, 2, 5}
		expectedOutputs := []int{3, 5}
		result := singleNumber(nums)

		assert.Equal(t, len(expectedOutputs), len(result))
		for _, expected := range expectedOutputs {
			assert.Contains(t, result, expected)
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{-1, 0}
		expectedOutputs := []int{-1, 0}
		result := singleNumber(nums)

		assert.Equal(t, len(expectedOutputs), len(result))
		for _, expected := range expectedOutputs {
			assert.Contains(t, result, expected)
		}
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{0, 1}
		expectedOutputs := []int{1, 0}
		result := singleNumber(nums)

		assert.Equal(t, len(expectedOutputs), len(result))
		for _, expected := range expectedOutputs {
			assert.Contains(t, result, expected)
		}
	})
}
