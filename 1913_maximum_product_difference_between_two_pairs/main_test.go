package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProductDifference(t *testing.T) {
	t.Run("should return 34", func(t *testing.T) {
		nums := []int{5, 6, 2, 7, 4}
		expected := 34
		assert.Equal(t, expected, maxProductDifference(nums))
	})

	t.Run("should return 64", func(t *testing.T) {
		nums := []int{4, 2, 5, 9, 7, 4, 8}
		expected := 64
		assert.Equal(t, expected, maxProductDifference(nums))
	})
}
