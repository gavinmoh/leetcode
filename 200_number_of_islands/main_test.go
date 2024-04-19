package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumIslands(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid := [][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}
		expected := 1

		assert.Equal(t, expected, numIslands(grid))
	})

	t.Run("test case 2", func(t *testing.T) {
		grid := [][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		}
		expected := 3

		assert.Equal(t, expected, numIslands(grid))
	})

	t.Run("test case 3", func(t *testing.T) {
		grid := [][]byte{
			{'1', '1', '1'},
			{'0', '1', '0'},
			{'1', '1', '1'},
		}
		expected := 1

		assert.Equal(t, expected, numIslands(grid))
	})
}
