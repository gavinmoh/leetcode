package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumImportance(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 5
		roads := [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}}
		expected := int64(43)

		assert.Equal(t, expected, maximumImportance(n, roads))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 5
		roads := [][]int{{0, 3}, {2, 4}, {1, 3}}
		expected := int64(20)

		assert.Equal(t, expected, maximumImportance(n, roads))
	})
}
