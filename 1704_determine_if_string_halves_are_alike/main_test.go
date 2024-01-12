package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHalvesAreAlike(t *testing.T) {
	t.Run("should return true", func(t *testing.T) {
		s := "book"
		assert.True(t, halvesAreAlike(s))
	})

	t.Run("should return false", func(t *testing.T) {
		s := "textbook"
		assert.False(t, halvesAreAlike(s))
	})
}
