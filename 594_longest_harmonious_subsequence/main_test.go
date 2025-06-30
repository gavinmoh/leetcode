package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLHS(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
		expected := 5

		assert.Equal(t, expected, findLHS(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		expected := 2

		assert.Equal(t, expected, findLHS(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{1, 1, 1, 1}
		expected := 0

		assert.Equal(t, expected, findLHS(nums))
	})

	t.Run("test case 4", func(t *testing.T) {
		nums := []int{-3, -1, -1, -1, -3, -2}
		expected := 4

		assert.Equal(t, expected, findLHS(nums))
	})
}
