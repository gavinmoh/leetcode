package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximalRectangle(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		matrix := [][]byte{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		}
		expected := 6

		assert.Equal(t, expected, maximalRectangle(matrix))
	})

	t.Run("test case 2", func(t *testing.T) {
		matrix := [][]byte{
			{'0'},
		}
		expected := 0

		assert.Equal(t, expected, maximalRectangle(matrix))
	})

	t.Run("test case 3", func(t *testing.T) {
		matrix := [][]byte{
			{'1'},
		}
		expected := 1

		assert.Equal(t, expected, maximalRectangle(matrix))
	})

	t.Run("test case 4", func(t *testing.T) {
		matrix := [][]byte{
			{'0', '1'},
			{'1', '0'},
		}
		expected := 1

		assert.Equal(t, expected, maximalRectangle(matrix))
	})
}
