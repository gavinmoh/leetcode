package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContinuousSubarrays(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{5, 4, 2, 4}
		expected := int64(8)

		assert.Equal(t, expected, continuousSubarrays(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2, 3}
		expected := int64(6)

		assert.Equal(t, expected, continuousSubarrays(nums))
	})
}
