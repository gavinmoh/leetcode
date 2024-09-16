package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinDifference(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		timePoints := []string{"23:59", "00:00"}
		expected := 1

		assert.Equal(t, expected, findMinDifference(timePoints))
	})

	t.Run("test case 2", func(t *testing.T) {
		timePoints := []string{"00:00", "23:59", "00:00"}
		expected := 0

		assert.Equal(t, expected, findMinDifference(timePoints))
	})
}
