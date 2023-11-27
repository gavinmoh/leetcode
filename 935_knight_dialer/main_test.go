package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnightDialer(t *testing.T) {
	t.Run("should return 10", func(t *testing.T) {
		n := 1
		expected := 10
		assert.Equal(t, expected, knightDialer(n))
	})

	t.Run("should return 20", func(t *testing.T) {
		n := 2
		expected := 20
		assert.Equal(t, expected, knightDialer(n))
	})

	t.Run("should return 136006598", func(t *testing.T) {
		n := 3131
		expected := 136006598
		assert.Equal(t, expected, knightDialer(n))
	})
}
