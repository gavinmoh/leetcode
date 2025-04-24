package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCompleteSubarrays(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 3, 1, 2, 2}
		expected := 4

		assert.Equal(t, expected, countCompleteSubarrays(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{5, 5, 5, 5}
		expected := 10

		assert.Equal(t, expected, countCompleteSubarrays(nums))
	})
}
