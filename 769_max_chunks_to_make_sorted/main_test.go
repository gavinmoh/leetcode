package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxChunksToSorted(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		arr := []int{4, 3, 2, 1, 0}
		expected := 1

		assert.Equal(t, expected, maxChunksToSorted(arr))
	})

	t.Run("test case 2", func(t *testing.T) {
		arr := []int{1, 0, 2, 3, 4}
		expected := 4

		assert.Equal(t, expected, maxChunksToSorted(arr))
	})
}
