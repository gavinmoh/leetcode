package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestNiceSubarray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 3, 8, 48, 10}
		expected := 3

		assert.Equal(t, expected, longestNiceSubarray(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{3, 1, 5, 11, 13}
		expected := 1

		assert.Equal(t, expected, longestNiceSubarray(nums))
	})
}
