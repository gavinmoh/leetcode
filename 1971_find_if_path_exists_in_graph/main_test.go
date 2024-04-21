package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPath(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 3
		edges := [][]int{{0, 1}, {1, 2}, {2, 0}}
		source := 0
		destination := 2

		assert.True(t, validPath(n, edges, source, destination))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 6
		edges := [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}
		source := 0
		destination := 5

		assert.False(t, validPath(n, edges, source, destination))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 50
		edges := [][]int{
			{31, 5}, {10, 46}, {19, 31}, {5, 1}, {31, 28},
			{28, 29}, {8, 26}, {13, 23}, {16, 34}, {30, 1},
			{16, 18}, {33, 46}, {27, 35}, {2, 25}, {49, 33},
			{44, 19}, {22, 26}, {30, 13}, {27, 12}, {8, 16},
			{42, 13}, {18, 3}, {21, 20}, {2, 17}, {5, 48},
			{41, 37}, {39, 37}, {2, 11}, {20, 26}, {19, 43},
			{45, 7}, {0, 21}, {44, 23}, {2, 39}, {27, 36},
			{41, 48}, {17, 42}, {40, 32}, {2, 28}, {35, 38},
			{3, 9}, {41, 30}, {5, 11}, {24, 22}, {39, 5},
			{40, 31}, {18, 35}, {23, 39}, {20, 24}, {45, 12},
		}
		source := 29
		destination := 46

		assert.False(t, validPath(n, edges, source, destination))
	})

	t.Run("test case 4", func(t *testing.T) {
		n := 1
		edges := [][]int{}
		source := 0
		destination := 0

		assert.True(t, validPath(n, edges, source, destination))
	})
}
