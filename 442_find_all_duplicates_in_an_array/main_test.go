package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicates(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
		expected := []int{2, 3}

		assert.Equal(t, expected, findDuplicates(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 1, 2}
		expected := []int{1}

		assert.Equal(t, expected, findDuplicates(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{1}
		expected := []int{}

		assert.Equal(t, expected, findDuplicates(nums))
	})
}
