package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroFilledSubarray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 3, 0, 0, 2, 0, 0, 4}
		expected := int64(6)

		assert.Equal(t, expected, zeroFilledSubarray(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{0, 0, 0, 2, 0, 0}
		expected := int64(9)

		assert.Equal(t, expected, zeroFilledSubarray(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{2, 10, 2019}
		expected := int64(0)

		assert.Equal(t, expected, zeroFilledSubarray(nums))
	})
}
