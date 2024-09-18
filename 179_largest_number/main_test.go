package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestNumber(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{10, 2}
		expected := "210"

		assert.Equal(t, expected, largestNumber(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{3, 30, 34, 5, 9}
		expected := "9534330"

		assert.Equal(t, expected, largestNumber(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{0, 0}
		expected := "0"

		assert.Equal(t, expected, largestNumber(nums))
	})
}
