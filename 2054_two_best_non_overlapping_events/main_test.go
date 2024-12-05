package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxTwoEvents(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		events := [][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}}
		expected := 4

		assert.Equal(t, expected, maxTwoEvents(events))
	})

	t.Run("test case 2", func(t *testing.T) {
		events := [][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}}
		expected := 5

		assert.Equal(t, expected, maxTwoEvents(events))
	})

	t.Run("test case 3", func(t *testing.T) {
		events := [][]int{{1, 5, 3}, {1, 5, 1}, {6, 6, 5}}
		expected := 8

		assert.Equal(t, expected, maxTwoEvents(events))
	})

	t.Run("test case 4", func(t *testing.T) {
		events := [][]int{
			{66, 97, 90}, {98, 98, 68}, {38, 49, 63},
			{91, 100, 42}, {92, 100, 22}, {1, 77, 50},
			{64, 72, 97},
		}
		expected := 165

		assert.Equal(t, expected, maxTwoEvents(events))
	})
}
