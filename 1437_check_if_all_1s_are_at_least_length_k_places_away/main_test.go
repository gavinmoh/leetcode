package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKLengthApart(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 0, 0, 0, 1, 0, 0, 1}
		k := 2

		assert.True(t, kLengthApart(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 0, 0, 1, 0, 1}
		k := 2

		assert.False(t, kLengthApart(nums, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{
			1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}
		k := 100

		assert.True(t, kLengthApart(nums, k))
	})
}
