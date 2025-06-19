package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartitionArray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{3, 6, 1, 2, 5}
		k := 2
		expected := 2

		assert.Equal(t, expected, partitionArray(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2, 3}
		k := 1
		expected := 2

		assert.Equal(t, expected, partitionArray(nums, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{2, 2, 4, 5}
		k := 0
		expected := 3

		assert.Equal(t, expected, partitionArray(nums, k))
	})
}
