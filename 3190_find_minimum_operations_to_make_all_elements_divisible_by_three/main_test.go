package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumOperations(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		expected := 3

		assert.Equal(t, expected, minimumOperations(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{3, 6, 9}
		expected := 0

		assert.Equal(t, expected, minimumOperations(nums))
	})
}
