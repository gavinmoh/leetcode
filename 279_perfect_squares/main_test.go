package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSquares(t *testing.T) {
	t.Run("should return 3", func(t *testing.T) {
		n := 12
		expected := 3

		assert.Equal(t, expected, numSquares(n))
	})

	t.Run("should return 2", func(t *testing.T) {
		n := 13
		expected := 2

		assert.Equal(t, expected, numSquares(n))
	})
}
