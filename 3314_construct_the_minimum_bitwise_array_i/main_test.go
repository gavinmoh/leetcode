package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinBitwiseArray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{2, 3, 5, 7}
		expected := []int{-1, 1, 4, 3}

		assert.Equal(t, expected, minBitwiseArray(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{11, 13, 31}
		expected := []int{9, 12, 15}

		assert.Equal(t, expected, minBitwiseArray(nums))
	})
}
