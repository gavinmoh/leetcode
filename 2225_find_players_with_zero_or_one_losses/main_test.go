package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWinners(t *testing.T) {
	t.Run("should return [[1,2,10],[4,5,7,8]]", func(t *testing.T) {
		matches := [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}
		expected := [][]int{{1, 2, 10}, {4, 5, 7, 8}}
		assert.Equal(t, expected, findWinners(matches))
	})

	t.Run("should return [[1,2,5,6],[]]", func(t *testing.T) {
		matches := [][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}}
		expected := [][]int{{1, 2, 5, 6}, {}}
		assert.Equal(t, expected, findWinners(matches))
	})
}
