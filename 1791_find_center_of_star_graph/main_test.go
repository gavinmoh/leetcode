package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCenter(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		edges := [][]int{{1, 2}, {2, 3}, {4, 2}}
		expected := 2

		assert.Equal(t, expected, findCenter(edges))
	})

	t.Run("test case 2", func(t *testing.T) {
		edges := [][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}
		expected := 1

		assert.Equal(t, expected, findCenter(edges))
	})
}
