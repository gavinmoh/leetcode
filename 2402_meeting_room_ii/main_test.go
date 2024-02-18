package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMostBooked(t *testing.T) {
	t.Run("should return 0", func(t *testing.T) {
		n := 2
		meetings := [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}

		expected := 0
		assert.Equal(t, expected, mostBooked(n, meetings))
	})

	t.Run("should return 1", func(t *testing.T) {
		n := 3
		meetings := [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}

		expected := 1
		assert.Equal(t, expected, mostBooked(n, meetings))
	})

	t.Run("should return 0", func(t *testing.T) {
		n := 4
		meetings := [][]int{{18, 19}, {3, 12}, {17, 19}, {2, 13}, {7, 10}}

		expected := 0
		assert.Equal(t, expected, mostBooked(n, meetings))
	})

	t.Run("should return 0", func(t *testing.T) {
		n := 4
		meetings := [][]int{{10, 11}, {13, 15}, {9, 19}, {0, 12}, {12, 20}}

		expected := 0
		assert.Equal(t, expected, mostBooked(n, meetings))
	})

	t.Run("should return 1", func(t *testing.T) {
		n := 4
		meetings := [][]int{{19, 20}, {14, 15}, {13, 14}, {11, 20}}

		expected := 1
		assert.Equal(t, expected, mostBooked(n, meetings))
	})
}
