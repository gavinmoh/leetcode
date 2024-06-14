package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinIncrementForUnique(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 2}
		expected := 1

		assert.Equal(t, expected, minIncrementForUnique(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{3, 2, 1, 2, 1, 7}
		expected := 6

		assert.Equal(t, expected, minIncrementForUnique(nums))
	})
}
