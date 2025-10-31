package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSneakyNumbers(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{0, 1, 1, 0}
		expected := []int{0, 1}

		assert.Equal(t, expected, getSneakyNumbers(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{0, 3, 2, 1, 3, 2}
		expected := []int{2, 3}

		assert.Equal(t, expected, getSneakyNumbers(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}
		expected := []int{4, 5}

		assert.Equal(t, expected, getSneakyNumbers(nums))
	})
}
