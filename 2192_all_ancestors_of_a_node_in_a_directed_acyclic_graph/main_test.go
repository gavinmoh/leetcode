package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAncestors(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 8
		edgeList := [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}}
		expected := [][]int{{}, {}, {}, {0, 1}, {0, 2}, {0, 1, 3}, {0, 1, 2, 3, 4}, {0, 1, 2, 3}}

		assert.Equal(t, expected, getAncestors(n, edgeList))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 5
		edgeList := [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}
		expected := [][]int{{}, {0}, {0, 1}, {0, 1, 2}, {0, 1, 2, 3}}

		assert.Equal(t, expected, getAncestors(n, edgeList))
	})
}
