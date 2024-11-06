package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountFairPairs(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{0, 1, 7, 4, 4, 5}
		lower := 3
		upper := 6
		expected := int64(6)

		assert.Equal(t, expected, countFairPairs(nums, lower, upper))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 7, 9, 2, 5}
		lower := 11
		upper := 11
		expected := int64(1)

		assert.Equal(t, expected, countFairPairs(nums, lower, upper))
	})
}
