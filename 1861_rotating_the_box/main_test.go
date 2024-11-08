package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateTheBox(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		box := [][]byte{{'#', '.', '#'}}
		expected := [][]byte{{'.'}, {'#'}, {'#'}}

		assert.Equal(t, expected, rotateTheBox(box))
	})

	t.Run("test case 2", func(t *testing.T) {
		box := [][]byte{
			{'#', '.', '*', '.'},
			{'#', '#', '*', '.'},
		}
		expected := [][]byte{
			{'#', '.'},
			{'#', '#'},
			{'*', '*'},
			{'.', '.'},
		}

		assert.Equal(t, expected, rotateTheBox(box))
	})

	t.Run("test case 3", func(t *testing.T) {
		box := [][]byte{
			{'#', '#', '*', '.', '*', '.'},
			{'#', '#', '#', '*', '.', '.'},
			{'#', '#', '#', '.', '#', '.'},
		}
		expected := [][]byte{
			{'.', '#', '#'},
			{'.', '#', '#'},
			{'#', '#', '*'},
			{'#', '*', '.'},
			{'#', '.', '*'},
			{'#', '.', '.'},
		}

		assert.Equal(t, expected, rotateTheBox(box))
	})
}
