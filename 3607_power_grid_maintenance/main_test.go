package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessQueries(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		c := 5
		connections := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
		queries := [][]int{{1, 3}, {2, 1}, {1, 1}, {2, 2}, {1, 2}}
		expected := []int{3, 2, 3}

		assert.Equal(t, expected, processQueries(c, connections, queries))
	})

	t.Run("test case 2", func(t *testing.T) {
		c := 3
		connections := [][]int{}
		queries := [][]int{{1, 1}, {2, 1}, {1, 1}}
		expected := []int{1, -1}

		assert.Equal(t, expected, processQueries(c, connections, queries))
	})

	t.Run("test case 3", func(t *testing.T) {
		c := 3
		connections := [][]int{{2, 3}, {1, 2}, {1, 3}}
		queries := [][]int{
			{1, 1}, {1, 2}, {1, 2}, {1, 3}, {1, 3},
			{1, 1}, {2, 3}, {1, 1}, {2, 2}, {2, 2},
			{1, 2}, {1, 3}, {2, 1}, {2, 1}, {1, 3},
			{2, 1}, {2, 3}, {1, 3}, {1, 3}, {2, 2},
			{1, 1}, {2, 2}, {1, 2}, {1, 1}, {1, 2},
			{1, 3}, {1, 2}, {1, 3}, {2, 2}, {2, 2},
			{2, 3}, {1, 3}, {1, 2}, {2, 3}, {1, 2},
			{2, 3}, {2, 3}, {2, 2}, {2, 2}, {1, 1},
			{2, 3}, {1, 1},
		}
		expected := []int{1, 2, 2, 3, 3, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}

		assert.Equal(t, expected, processQueries(c, connections, queries))
	})

	t.Run("test case 4", func(t *testing.T) {
		c := 3
		connections := [][]int{{3, 2}, {1, 3}, {2, 1}}
		queries := [][]int{
			{2, 2}, {1, 2}, {1, 2}, {1, 3}, {1, 1},
			{1, 3}, {1, 1}, {1, 1}, {2, 1}, {1, 1},
			{2, 3}, {2, 3}, {2, 3}, {2, 1}, {2, 1},
			{2, 1}, {1, 1}, {1, 1}, {1, 2}, {1, 2},
			{2, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1},
		}
		expected := []int{1, 1, 3, 1, 3, 1, 1, 3, -1, -1, -1, -1, -1, -1}

		assert.Equal(t, expected, processQueries(c, connections, queries))
	})
}
