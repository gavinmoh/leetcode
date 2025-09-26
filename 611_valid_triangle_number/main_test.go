package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriangleNumber(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{2, 2, 3, 4}
		expected := 3

		assert.Equal(t, expected, triangleNumber(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{4, 2, 3, 4}
		expected := 4

		assert.Equal(t, expected, triangleNumber(nums))
	})
}
