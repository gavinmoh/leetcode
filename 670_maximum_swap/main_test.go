package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumSwap(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		num := 2736
		expected := 7236

		assert.Equal(t, expected, maximumSwap(num))
	})

	t.Run("test case 2", func(t *testing.T) {
		num := 9973
		expected := 9973

		assert.Equal(t, expected, maximumSwap(num))
	})

	t.Run("test case 3", func(t *testing.T) {
		num := 98368
		expected := 98863

		assert.Equal(t, expected, maximumSwap(num))
	})

	t.Run("test case 4", func(t *testing.T) {
		num := 115
		expected := 511

		assert.Equal(t, expected, maximumSwap(num))
	})

	t.Run("test case 5", func(t *testing.T) {
		num := 10909091
		expected := 90909011

		assert.Equal(t, expected, maximumSwap(num))
	})
}
