package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumPairRemoval(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{5, 2, 3, 1}
		expected := 2

		assert.Equal(t, expected, minimumPairRemoval(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2, 2}
		expected := 0

		assert.Equal(t, expected, minimumPairRemoval(nums))
	})
}
