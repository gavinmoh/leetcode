package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpecialArray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{3, 5}
		expected := 2

		assert.Equal(t, expected, specialArray(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{0, 0}
		expected := -1

		assert.Equal(t, expected, specialArray(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{0, 4, 3, 0, 4}
		expected := 3

		assert.Equal(t, expected, specialArray(nums))
	})
}
