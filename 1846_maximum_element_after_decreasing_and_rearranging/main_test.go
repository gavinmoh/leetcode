package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumElementAfterDecrementingAndRearranging(t *testing.T) {
	t.Run("should return 2", func(t *testing.T) {
		arr := []int{2, 2, 1, 2, 1}
		expected := 2
		assert.Equal(t, expected, maximumElementAfterDecrementingAndRearranging(arr))
	})

	t.Run("should return 3", func(t *testing.T) {
		arr := []int{100, 1, 1000}
		expected := 3
		assert.Equal(t, expected, maximumElementAfterDecrementingAndRearranging(arr))
	})

	t.Run("should return 5", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		expected := 5
		assert.Equal(t, expected, maximumElementAfterDecrementingAndRearranging(arr))
	})
}
