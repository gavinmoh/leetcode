package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPivotInteger(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 8
		expected := 6

		assert.Equal(t, expected, pivotInteger(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 1
		expected := 1

		assert.Equal(t, expected, pivotInteger(n))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 4
		expected := -1

		assert.Equal(t, expected, pivotInteger(n))
	})
}
