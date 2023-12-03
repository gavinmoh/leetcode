package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinTimeToVisitAllPoints(t *testing.T) {
	t.Run("should return 7", func(t *testing.T) {
		points := [][]int{{1, 1}, {3, 4}, {-1, 0}}
		expected := 7
		assert.Equal(t, expected, minTimeToVisitAllPoints(points))
	})

	t.Run("should return 5", func(t *testing.T) {
		points := [][]int{{3, 2}, {-2, 2}}
		expected := 5
		assert.Equal(t, expected, minTimeToVisitAllPoints(points))
	})

}
