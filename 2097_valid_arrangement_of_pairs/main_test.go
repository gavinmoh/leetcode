package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidArrangement(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		pairs := [][]int{{5, 1}, {4, 5}, {11, 9}, {9, 4}}
		expected := [][]int{{11, 9}, {9, 4}, {4, 5}, {5, 1}}

		assert.Equal(t, expected, validArrangement(pairs))
	})

	t.Run("test case 2", func(t *testing.T) {
		pairs := [][]int{{1, 3}, {3, 2}, {2, 1}}
		expected := [][]int{{1, 3}, {3, 2}, {2, 1}}

		assert.Equal(t, expected, validArrangement(pairs))
	})

	t.Run("test case 3", func(t *testing.T) {
		pairs := [][]int{{1, 2}, {1, 3}, {2, 1}}
		expected := [][]int{{1, 2}, {2, 1}, {1, 3}}

		assert.Equal(t, expected, validArrangement(pairs))
	})
}
