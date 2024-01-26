package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPaths(t *testing.T) {
	t.Run("should return 6", func(t *testing.T) {
		m := 2
		n := 2
		maxMove := 2
		startRow := 0
		startColumn := 0
		expected := 6

		assert.Equal(t, expected, findPaths(m, n, maxMove, startRow, startColumn))
	})

	t.Run("should return 12", func(t *testing.T) {
		m := 1
		n := 3
		maxMove := 3
		startRow := 0
		startColumn := 1
		expected := 12

		assert.Equal(t, expected, findPaths(m, n, maxMove, startRow, startColumn))
	})

	t.Run("should return 914783380", func(t *testing.T) {
		m := 8
		n := 50
		maxMove := 23
		startRow := 5
		startColumn := 26
		expected := 914783380

		assert.Equal(t, expected, findPaths(m, n, maxMove, startRow, startColumn))
	})
}
