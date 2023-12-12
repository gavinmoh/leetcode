package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProduct(t *testing.T) {
	t.Run("should return 12", func(t *testing.T) {
		nums := []int{3, 4, 5, 2}
		expected := 12
		assert.Equal(t, expected, maxProduct(nums))
	})

	t.Run("should return 16", func(t *testing.T) {
		nums := []int{1, 5, 4, 5}
		expected := 16
		assert.Equal(t, expected, maxProduct(nums))
	})

	t.Run("should return 12", func(t *testing.T) {
		nums := []int{3, 7}
		expected := 12
		assert.Equal(t, expected, maxProduct(nums))
	})
}
