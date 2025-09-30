package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriangularSum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		expected := 8

		assert.Equal(t, expected, triangularSum(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{5}
		expected := 5

		assert.Equal(t, expected, triangularSum(nums))
	})
}
