package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubarraysDivByK(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{4, 5, 0, -2, -3, 1}
		k := 5
		expected := 7

		assert.Equal(t, expected, subarraysDivByK(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{5}
		k := 9
		expected := 0

		assert.Equal(t, expected, subarraysDivByK(nums, k))
	})
}
