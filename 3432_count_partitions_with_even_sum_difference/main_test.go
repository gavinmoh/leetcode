package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPartitions(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{10, 10, 3, 7, 6}
		expected := 4

		assert.Equal(t, expected, countPartitions(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2, 2}
		expected := 0

		assert.Equal(t, expected, countPartitions(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{2, 4, 6, 8}
		expected := 3

		assert.Equal(t, expected, countPartitions(nums))
	})
}
