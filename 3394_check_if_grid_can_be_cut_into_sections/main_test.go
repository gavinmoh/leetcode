package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckValidCuts(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 5
		rectangles := [][]int{{1, 0, 5, 2}, {0, 2, 2, 4}, {3, 2, 5, 3}, {0, 4, 4, 5}}

		assert.True(t, checkValidCuts(n, rectangles))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 4
		rectangles := [][]int{{0, 0, 1, 1}, {2, 0, 3, 4}, {0, 2, 2, 3}, {3, 0, 4, 3}}

		assert.True(t, checkValidCuts(n, rectangles))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 4
		rectangles := [][]int{{0, 2, 2, 4}, {1, 0, 3, 2}, {2, 2, 3, 4}, {3, 0, 4, 2}, {3, 2, 4, 4}}

		assert.False(t, checkValidCuts(n, rectangles))
	})
}
