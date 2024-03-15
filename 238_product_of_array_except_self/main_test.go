package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		expected := []int{24, 12, 8, 6}

		assert.Equal(t, expected, productExceptSelf(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{-1, 1, 0, -3, 3}
		expected := []int{0, 0, 9, 0, 0}

		assert.Equal(t, expected, productExceptSelf(nums))
	})
}
