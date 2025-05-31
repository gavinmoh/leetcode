package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnakesAndLadders(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		board := [][]int{
			{-1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, 35, -1, -1, 13, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, 15, -1, -1, -1, -1},
		}
		expected := 4

		assert.Equal(t, expected, snakesAndLadders(board))
	})

	t.Run("test case 2", func(t *testing.T) {
		board := [][]int{
			{-1, -1},
			{-1, 3},
		}
		expected := 1

		assert.Equal(t, expected, snakesAndLadders(board))
	})

	t.Run("test case 3", func(t *testing.T) {
		board := [][]int{
			{1, 1, -1},
			{1, 1, 1},
			{-1, 1, 1},
		}
		expected := -1

		assert.Equal(t, expected, snakesAndLadders(board))
	})
}
