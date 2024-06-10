package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeightChecker(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		heights := []int{1, 1, 4, 2, 1, 3}
		expected := 3

		assert.Equal(t, expected, heightChecker(heights))
	})

	t.Run("test case 2", func(t *testing.T) {
		heights := []int{5, 1, 2, 3, 4}
		expected := 5

		assert.Equal(t, expected, heightChecker(heights))
	})

	t.Run("test case 3", func(t *testing.T) {
		heights := []int{1, 2, 3, 4, 5}
		expected := 0

		assert.Equal(t, expected, heightChecker(heights))
	})
}
