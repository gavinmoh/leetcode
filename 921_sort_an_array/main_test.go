package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortArray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{5, 2, 3, 1}
		expected := []int{1, 2, 3, 5}

		assert.Equal(t, expected, sortArray(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{5, 1, 1, 2, 0, 0}
		expected := []int{0, 0, 1, 1, 2, 5}

		assert.Equal(t, expected, sortArray(nums))
	})
}
