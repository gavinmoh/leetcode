package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxWidthRamp(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{6, 0, 8, 2, 1, 5}
		expected := 4

		assert.Equal(t, expected, maxWidthRamp(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}
		expected := 7

		assert.Equal(t, expected, maxWidthRamp(nums))
	})
}
