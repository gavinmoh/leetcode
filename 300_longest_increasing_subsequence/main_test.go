package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLIS(t *testing.T) {
	t.Run("should return 4", func(t *testing.T) {
		nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
		expected := 4
		assert.Equal(t, expected, lengthOfLIS(nums))
	})

	t.Run("should return 4", func(t *testing.T) {
		nums := []int{0, 1, 0, 3, 2, 3}
		expected := 4
		assert.Equal(t, expected, lengthOfLIS(nums))
	})

	t.Run("should return 1", func(t *testing.T) {
		nums := []int{7, 7, 7, 7, 7, 7, 7}
		expected := 1
		assert.Equal(t, expected, lengthOfLIS(nums))
	})
}
