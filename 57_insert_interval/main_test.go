package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		intervals := [][]int{{1, 3}, {6, 9}}
		newInterval := []int{2, 5}
		expected := [][]int{{1, 5}, {6, 9}}

		assert.Equal(t, expected, insert(intervals, newInterval))
	})

	t.Run("test case 2", func(t *testing.T) {
		intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
		newInterval := []int{4, 8}
		expected := [][]int{{1, 2}, {3, 10}, {12, 16}}

		assert.Equal(t, expected, insert(intervals, newInterval))
	})
}
