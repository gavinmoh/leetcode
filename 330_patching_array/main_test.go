package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinPatches(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 3}
		n := 6
		expected := 1

		assert.Equal(t, expected, minPatches(nums, n))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 5, 10}
		n := 20
		expected := 2

		assert.Equal(t, expected, minPatches(nums, n))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{1, 2, 2}
		n := 5
		expected := 0

		assert.Equal(t, expected, minPatches(nums, n))
	})
}
