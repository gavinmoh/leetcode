package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreaOfMaxDiagonal(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		dimensions := [][]int{{9, 3}, {8, 6}}
		expected := 48

		assert.Equal(t, expected, areaOfMaxDiagonal(dimensions))
	})

	t.Run("test case 2", func(t *testing.T) {
		dimensions := [][]int{{4, 3}, {3, 4}}
		expected := 12

		assert.Equal(t, expected, areaOfMaxDiagonal(dimensions))
	})
}
