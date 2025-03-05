package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColoredCells(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 1
		expected := int64(1)

		assert.Equal(t, expected, coloredCells(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 2
		expected := int64(5)

		assert.Equal(t, expected, coloredCells(n))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 3
		expected := int64(13)

		assert.Equal(t, expected, coloredCells(n))
	})

	t.Run("test case 4", func(t *testing.T) {
		n := 4
		expected := int64(25)

		assert.Equal(t, expected, coloredCells(n))
	})
}
