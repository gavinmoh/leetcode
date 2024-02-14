package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRearrangeArray(t *testing.T) {
	t.Run("should return [3, -2, 1, -5, 2, -4]", func(t *testing.T) {
		nums := []int{3, 1, -2, -5, 2, -4}
		expected := []int{3, -2, 1, -5, 2, -4}

		assert.Equal(t, expected, rearrangeArray(nums))
	})

	t.Run("should return [1, -1]", func(t *testing.T) {
		nums := []int{-1, 1}
		expected := []int{1, -1}

		assert.Equal(t, expected, rearrangeArray(nums))
	})
}
