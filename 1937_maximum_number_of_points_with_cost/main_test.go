package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxPoints(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		points := [][]int{{1, 2, 3}, {1, 5, 1}, {3, 1, 1}}
		expected := int64(9)

		assert.Equal(t, expected, maxPoints(points))
	})

	t.Run("test case 2", func(t *testing.T) {
		points := [][]int{{1, 5}, {2, 3}, {4, 2}}
		expected := int64(11)

		assert.Equal(t, expected, maxPoints(points))
	})
}
