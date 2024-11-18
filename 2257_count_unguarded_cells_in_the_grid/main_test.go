package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountUnguarded(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		m := 4
		n := 6
		guards := [][]int{{0, 0}, {1, 1}, {2, 3}}
		walls := [][]int{{0, 1}, {2, 2}, {1, 4}}
		expected := 7

		assert.Equal(t, expected, countUnguarded(m, n, guards, walls))
	})

	t.Run("test case 2", func(t *testing.T) {
		m := 3
		n := 3
		guards := [][]int{{1, 1}}
		walls := [][]int{{0, 1}, {1, 0}, {2, 1}, {1, 2}}
		expected := 4

		assert.Equal(t, expected, countUnguarded(m, n, guards, walls))
	})
}
