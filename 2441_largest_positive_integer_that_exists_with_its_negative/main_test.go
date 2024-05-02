package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxK(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{-1, 2, -3, 3}
		expected := 3

		assert.Equal(t, expected, findMaxK(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{-1, 10, 6, 7, -7, 1}
		expected := 7

		assert.Equal(t, expected, findMaxK(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{-10, 8, 6, 7, -2, -3}
		expected := -1

		assert.Equal(t, expected, findMaxK(nums))
	})
}
