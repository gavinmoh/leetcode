package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBeautifulSubsets(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{2, 4, 6}
		k := 2
		expected := 4

		assert.Equal(t, expected, beautifulSubsets(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1}
		k := 1
		expected := 1

		assert.Equal(t, expected, beautifulSubsets(nums, k))
	})
}
