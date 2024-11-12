package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLengthOfShortestSubarray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		arr := []int{1, 2, 3, 10, 4, 2, 3, 5}
		expected := 3

		assert.Equal(t, expected, findLengthOfShortestSubarray(arr))
	})

	t.Run("test case 2", func(t *testing.T) {
		arr := []int{5, 4, 3, 2, 1}
		expected := 4

		assert.Equal(t, expected, findLengthOfShortestSubarray(arr))
	})

	t.Run("test case 3", func(t *testing.T) {
		arr := []int{1, 2, 3}
		expected := 0

		assert.Equal(t, expected, findLengthOfShortestSubarray(arr))
	})
}
