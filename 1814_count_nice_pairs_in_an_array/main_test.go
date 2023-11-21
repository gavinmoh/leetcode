package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountNicePairs(t *testing.T) {
	t.Run("should return 2", func(t *testing.T) {
		nums := []int{42, 11, 1, 97}
		expected := 2

		assert.Equal(t, expected, countNicePairs(nums))
	})

	t.Run("should return 4", func(t *testing.T) {
		nums := []int{13, 10, 35, 24, 76}
		expected := 4

		assert.Equal(t, expected, countNicePairs(nums))
	})
}
