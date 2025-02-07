package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryResults(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		limit := 4
		queries := [][]int{{1, 4}, {2, 5}, {1, 3}, {3, 4}}
		expected := []int{1, 2, 2, 3}

		assert.Equal(t, expected, queryResults(limit, queries))
	})

	t.Run("test case 2", func(t *testing.T) {
		limit := 4
		queries := [][]int{{0, 1}, {1, 2}, {2, 2}, {3, 4}, {4, 5}}
		expected := []int{1, 2, 2, 3, 4}

		assert.Equal(t, expected, queryResults(limit, queries))
	})
}
