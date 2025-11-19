package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFinalValue(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{5, 3, 6, 1, 12}
		original := 3
		expected := 24

		assert.Equal(t, expected, findFinalValue(nums, original))
	})

	t.Run("test case 1", func(t *testing.T) {
		nums := []int{2, 7, 9}
		original := 4
		expected := 4

		assert.Equal(t, expected, findFinalValue(nums, original))
	})
}
