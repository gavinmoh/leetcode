package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumBeauty(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		items := [][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}
		queries := []int{1, 2, 3, 4, 5, 6}
		expected := []int{2, 4, 5, 5, 6, 6}

		assert.Equal(t, expected, maximumBeauty(items, queries))
	})

	t.Run("test case 2", func(t *testing.T) {
		items := [][]int{{1, 2}, {1, 2}, {1, 3}, {1, 4}}
		queries := []int{1}
		expected := []int{4}

		assert.Equal(t, expected, maximumBeauty(items, queries))
	})

	t.Run("test case 3", func(t *testing.T) {
		items := [][]int{{10, 1000}}
		queries := []int{5}
		expected := []int{0}

		assert.Equal(t, expected, maximumBeauty(items, queries))
	})
}
