package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFinalState(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{2, 1, 3, 5, 6}
		k := 5
		multiplier := 2
		expected := []int{8, 4, 6, 5, 6}

		assert.Equal(t, expected, getFinalState(nums, k, multiplier))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2}
		k := 3
		multiplier := 4
		expected := []int{16, 8}

		assert.Equal(t, expected, getFinalState(nums, k, multiplier))
	})
}
