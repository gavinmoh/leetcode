package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindClosest(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		x := 2
		y := 7
		z := 4
		expected := 1

		assert.Equal(t, expected, findClosest(x, y, z))
	})

	t.Run("test case 2", func(t *testing.T) {
		x := 2
		y := 5
		z := 6
		expected := 2

		assert.Equal(t, expected, findClosest(x, y, z))
	})

	t.Run("test case 3", func(t *testing.T) {
		x := 1
		y := 5
		z := 3
		expected := 0

		assert.Equal(t, expected, findClosest(x, y, z))
	})
}
