package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayRankTransform(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		arr := []int{40, 10, 20, 30}
		expected := []int{4, 1, 2, 3}

		assert.Equal(t, expected, arrayRankTransform(arr))
	})

	t.Run("test case 2", func(t *testing.T) {
		arr := []int{100, 100, 100}
		expected := []int{1, 1, 1}

		assert.Equal(t, expected, arrayRankTransform(arr))
	})

	t.Run("test case 3", func(t *testing.T) {
		arr := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}
		expected := []int{5, 3, 4, 2, 8, 6, 7, 1, 3}

		assert.Equal(t, expected, arrayRankTransform(arr))
	})
}
