package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegionsBySlashes(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid := []string{" /", "/ "}
		expected := 2

		assert.Equal(t, expected, regionsBySlashes(grid))
	})

	t.Run("test case 2", func(t *testing.T) {
		grid := []string{" /", "  "}
		expected := 1

		assert.Equal(t, expected, regionsBySlashes(grid))
	})

	t.Run("test case 3", func(t *testing.T) {
		grid := []string{"/\\", "\\/"}
		expected := 5

		assert.Equal(t, expected, regionsBySlashes(grid))
	})
}
