package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		arr1 := []int{1, 10, 100}
		arr2 := []int{1000}
		expected := 3

		assert.Equal(t, expected, longestCommonPrefix(arr1, arr2))
	})

	t.Run("test case 2", func(t *testing.T) {
		arr1 := []int{1, 2, 3}
		arr2 := []int{4, 4, 4}
		expected := 0

		assert.Equal(t, expected, longestCommonPrefix(arr1, arr2))
	})
}
