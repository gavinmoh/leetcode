package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasSameDigits(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "3902"

		assert.True(t, hasSameDigits(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "34789"

		assert.False(t, hasSameDigits(s))
	})
}
