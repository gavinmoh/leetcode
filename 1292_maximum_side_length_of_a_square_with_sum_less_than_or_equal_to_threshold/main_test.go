package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSideLength(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		mat := [][]int{{1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}}
		threshold := 4
		expected := 2

		assert.Equal(t, expected, maxSideLength(mat, threshold))
	})

	t.Run("test case 2", func(t *testing.T) {
		mat := [][]int{{2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}}
		threshold := 1
		expected := 0

		assert.Equal(t, expected, maxSideLength(mat, threshold))
	})

	t.Run("test case 3", func(t *testing.T) {
		mat := [][]int{{18, 70}, {61, 1}, {25, 85}, {14, 40}, {11, 96}, {97, 96}, {63, 45}}
		threshold := 40184
		expected := 2

		assert.Equal(t, expected, maxSideLength(mat, threshold))
	})
}
