package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountMaxOrSubsets(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{3, 1}
		expected := 2

		assert.Equal(t, expected, countMaxOrSubsets(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 2, 2}
		expected := 7

		assert.Equal(t, expected, countMaxOrSubsets(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{3, 2, 1, 5}
		expected := 6

		assert.Equal(t, expected, countMaxOrSubsets(nums))
	})
}
