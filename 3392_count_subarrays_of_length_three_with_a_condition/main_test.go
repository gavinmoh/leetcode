package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSubarrays(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 1, 4, 1}
		expected := 1

		assert.Equal(t, expected, countSubarrays(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 1, 1}
		expected := 0

		assert.Equal(t, expected, countSubarrays(nums))
	})
}
