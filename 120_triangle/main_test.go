package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumTotal(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
		expected := 11

		assert.Equal(t, expected, minimumTotal(triangle))
	})

	t.Run("test case 2", func(t *testing.T) {
		triangle := [][]int{{-10}}
		expected := -10

		assert.Equal(t, expected, minimumTotal(triangle))
	})
}
