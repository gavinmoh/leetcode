package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumDifference(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{90}
		k := 1
		expected := 0

		assert.Equal(t, expected, minimumDifference(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{9, 4, 1, 7}
		k := 2
		expected := 2

		assert.Equal(t, expected, minimumDifference(nums, k))
	})
}
