package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFarmLand(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		land := [][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}
		expected := [][]int{{0, 0, 0, 0}, {1, 1, 2, 2}}

		assert.Equal(t, expected, findFarmland(land))
	})

	t.Run("test case 2", func(t *testing.T) {
		land := [][]int{{1, 1}, {1, 1}}
		expected := [][]int{{0, 0, 1, 1}}

		assert.Equal(t, expected, findFarmland(land))
	})

	t.Run("test case 3", func(t *testing.T) {
		land := [][]int{{0}}
		expected := [][]int{}

		assert.Equal(t, expected, findFarmland(land))
	})
}
