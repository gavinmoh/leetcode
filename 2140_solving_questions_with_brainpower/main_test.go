package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMostPoints(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		questions := [][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}
		expected := int64(5)

		assert.Equal(t, expected, mostPoints(questions))
	})

	t.Run("test case 2", func(t *testing.T) {
		questions := [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}
		expected := int64(7)

		assert.Equal(t, expected, mostPoints(questions))
	})
}
