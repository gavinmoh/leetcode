package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortByBits(t *testing.T) {
	t.Run("should return [0, 1, 2, 4, 8, 3, 5, 6, 7]", func(t *testing.T) {
		arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
		expected := []int{0, 1, 2, 4, 8, 3, 5, 6, 7}
		assert.Equal(t, expected, sortByBits(arr))
	})

	t.Run("should return [1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024]", func(t *testing.T) {
		arr := []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}
		expected := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
		assert.Equal(t, expected, sortByBits(arr))
	})
}
