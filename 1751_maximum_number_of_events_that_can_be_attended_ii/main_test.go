package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxValue(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		events := [][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 1}}
		k := 2
		expected := 7

		assert.Equal(t, expected, maxValue(events, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		events := [][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 10}}
		k := 2
		expected := 10

		assert.Equal(t, expected, maxValue(events, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		events := [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}, {4, 4, 4}}
		k := 3
		expected := 9

		assert.Equal(t, expected, maxValue(events, k))
	})
}
