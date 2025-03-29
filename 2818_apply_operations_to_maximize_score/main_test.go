package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumScore(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{8, 3, 9, 3, 8}
		k := 2
		expected := 81

		assert.Equal(t, expected, maximumScore(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{19, 12, 14, 6, 10, 18}
		k := 3
		expected := 4788

		assert.Equal(t, expected, maximumScore(nums, k))
	})
}
