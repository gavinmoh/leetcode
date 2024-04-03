package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExist(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
		word := "ABCCED"

		assert.True(t, exist(board, word))
	})

	t.Run("test case 2", func(t *testing.T) {
		board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
		word := "SEE"

		assert.True(t, exist(board, word))
	})

	t.Run("test case 3", func(t *testing.T) {
		board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
		word := "ABCB"

		assert.False(t, exist(board, word))
	})

	t.Run("test case 4", func(t *testing.T) {
		board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}
		word := "ABCESEEEFS"

		assert.True(t, exist(board, word))
	})

	t.Run("test case 5", func(t *testing.T) {
		board := [][]byte{{'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a'}}
		word := "aaaaaaaaaaaaa"

		assert.False(t, exist(board, word))
	})
}
