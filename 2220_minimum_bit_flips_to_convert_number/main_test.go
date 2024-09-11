package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinBitFlips(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		start := 10
		goal := 7
		expected := 3

		assert.Equal(t, expected, minBitFlips(start, goal))
	})

	t.Run("test case 2", func(t *testing.T) {
		start := 3
		goal := 4
		expected := 3

		assert.Equal(t, expected, minBitFlips(start, goal))
	})
}
