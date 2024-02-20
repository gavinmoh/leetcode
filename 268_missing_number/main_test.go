package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingNumber(t *testing.T) {
	t.Run("should return 2", func(t *testing.T) {
		nums := []int{3, 0, 1}
		expected := 2

		assert.Equal(t, expected, missingNumber(nums))
	})

	t.Run("should return 2", func(t *testing.T) {
		nums := []int{0, 1}
		expected := 2

		assert.Equal(t, expected, missingNumber(nums))
	})

	t.Run("should return 8", func(t *testing.T) {
		nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
		expected := 8

		assert.Equal(t, expected, missingNumber(nums))
	})

}
