package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLenLongestFibSubseq(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
		expected := 5

		assert.Equal(t, expected, lenLongestFibSubseq(arr))
	})

	t.Run("test case 2", func(t *testing.T) {
		arr := []int{1, 3, 7, 11, 12, 14, 18}
		expected := 3

		assert.Equal(t, expected, lenLongestFibSubseq(arr))
	})
}
