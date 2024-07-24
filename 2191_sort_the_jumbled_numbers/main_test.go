package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortJumbled(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		mapping := []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}
		nums := []int{991, 338, 38}
		expected := []int{338, 38, 991}

		assert.Equal(t, expected, sortJumbled(mapping, nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		mapping := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		nums := []int{789, 456, 123}
		expected := []int{123, 456, 789}

		assert.Equal(t, expected, sortJumbled(mapping, nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		mapping := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		nums := []int{0, 999999999}
		expected := []int{0, 999999999}

		assert.Equal(t, expected, sortJumbled(mapping, nums))
	})

	t.Run("test case 4", func(t *testing.T) {
		mapping := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
		nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		expected := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

		assert.Equal(t, expected, sortJumbled(mapping, nums))
	})
}
