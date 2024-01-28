package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubarraySum(t *testing.T) {
	t.Run("should return 2", func(t *testing.T) {
		nums := []int{1, 1, 1}
		k := 2
		expected := 2
		assert.Equal(t, expected, subarraySum(nums, k))
	})

	t.Run("should return 2", func(t *testing.T) {
		nums := []int{1, 2, 3}
		k := 3
		expected := 2
		assert.Equal(t, expected, subarraySum(nums, k))
	})
}
