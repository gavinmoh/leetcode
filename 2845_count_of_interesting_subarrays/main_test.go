package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountInterestingSubarrays(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{3, 2, 4}
		modulo := 2
		k := 1
		expected := int64(3)

		assert.Equal(t, expected, countInterestingSubarrays(nums, modulo, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{3, 1, 9, 6}
		modulo := 3
		k := 0
		expected := int64(2)

		assert.Equal(t, expected, countInterestingSubarrays(nums, modulo, k))
	})

}
