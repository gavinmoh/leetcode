package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasAlternatingBits(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 5

		assert.True(t, hasAlternatingBits(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 7

		assert.False(t, hasAlternatingBits(n))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 11

		assert.False(t, hasAlternatingBits(n))
	})
}
