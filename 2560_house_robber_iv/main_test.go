package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCapability(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{2, 3, 5, 9}
		k := 2
		expected := 5

		assert.Equal(t, expected, minCapability(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 7, 9, 3, 1}
		k := 2
		expected := 2

		assert.Equal(t, expected, minCapability(nums, k))
	})
}
