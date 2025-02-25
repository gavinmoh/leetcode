package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumOfSubarrays(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		arr := []int{1, 3, 5}
		expected := 4

		assert.Equal(t, expected, numOfSubarrays(arr))
	})

	t.Run("test case 2", func(t *testing.T) {
		arr := []int{2, 4, 6}
		expected := 0

		assert.Equal(t, expected, numOfSubarrays(arr))
	})

	t.Run("test case 3", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 6, 7}
		expected := 16

		assert.Equal(t, expected, numOfSubarrays(arr))
	})
}
