package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCheapestPrice(t *testing.T) {
	t.Run("should return 700", func(t *testing.T) {
		n := 4
		flights := [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}
		src := 0
		dst := 3
		k := 1
		expected := 700

		assert.Equal(t, expected, findCheapestPrice(n, flights, src, dst, k))
	})

	t.Run("should return 200", func(t *testing.T) {
		n := 3
		flights := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
		src := 0
		dst := 2
		k := 1
		expected := 200

		assert.Equal(t, expected, findCheapestPrice(n, flights, src, dst, k))
	})

	t.Run("should return 500", func(t *testing.T) {
		n := 3
		flights := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
		src := 0
		dst := 2
		k := 0
		expected := 500

		assert.Equal(t, expected, findCheapestPrice(n, flights, src, dst, k))
	})

	t.Run("should return 5", func(t *testing.T) {
		n := 5
		flights := [][]int{{1, 2, 10}, {2, 0, 7}, {1, 3, 8}, {4, 0, 10}, {3, 4, 2}, {4, 2, 10}, {0, 3, 3}, {3, 1, 6}, {2, 4, 5}}
		src := 0
		dst := 4
		k := 1
		expected := 5

		assert.Equal(t, expected, findCheapestPrice(n, flights, src, dst, k))
	})
}
