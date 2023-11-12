package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumBusesToDestination(t *testing.T) {
	t.Run("should return 2", func(t *testing.T) {
		routes := [][]int{{1, 2, 7}, {3, 6, 7}}
		source := 1
		target := 6
		expected := 2

		assert.Equal(t, expected, numBusesToDestination(routes, source, target))
	})

	t.Run("should return -1", func(t *testing.T) {
		routes := [][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}
		source := 15
		target := 12
		expected := -1

		assert.Equal(t, expected, numBusesToDestination(routes, source, target))
	})

}
