package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfAlternatingGroups(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		colors := []int{0, 1, 0, 1, 0}
		k := 3
		expected := 3

		assert.Equal(t, expected, numberOfAlternatingGroups(colors, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		colors := []int{0, 1, 0, 0, 1, 0, 1}
		k := 6
		expected := 2

		assert.Equal(t, expected, numberOfAlternatingGroups(colors, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		colors := []int{1, 1, 0, 1}
		k := 4
		expected := 0

		assert.Equal(t, expected, numberOfAlternatingGroups(colors, k))
	})
}
