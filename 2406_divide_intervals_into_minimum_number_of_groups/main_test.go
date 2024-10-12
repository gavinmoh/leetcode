package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinGroups(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		intervals := [][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}}
		expected := 3

		assert.Equal(t, expected, minGroups(intervals))
	})

	t.Run("test case 2", func(t *testing.T) {
		intervals := [][]int{{1, 3}, {5, 6}, {8, 10}, {11, 13}}
		expected := 1

		assert.Equal(t, expected, minGroups(intervals))
	})
}
