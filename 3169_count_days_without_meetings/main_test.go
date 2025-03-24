package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDays(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		days := 10
		meetings := [][]int{{5, 7}, {1, 3}, {9, 10}}
		expected := 2

		assert.Equal(t, expected, countDays(days, meetings))
	})

	t.Run("test case 2", func(t *testing.T) {
		days := 5
		meetings := [][]int{{2, 4}, {1, 3}}
		expected := 1

		assert.Equal(t, expected, countDays(days, meetings))
	})

	t.Run("test case 3", func(t *testing.T) {
		days := 6
		meetings := [][]int{{1, 6}}
		expected := 0

		assert.Equal(t, expected, countDays(days, meetings))
	})

	t.Run("test case 4", func(t *testing.T) {
		days := 8
		meetings := [][]int{{3, 4}, {4, 8}, {2, 5}, {3, 8}}
		expected := 1

		assert.Equal(t, expected, countDays(days, meetings))
	})

	t.Run("test case 5", func(t *testing.T) {
		days := 57
		meetings := [][]int{{3, 49}, {23, 44}, {21, 56}, {26, 55}, {23, 52}, {2, 9}, {1, 48}, {3, 31}}
		expected := 1

		assert.Equal(t, expected, countDays(days, meetings))
	})
}
