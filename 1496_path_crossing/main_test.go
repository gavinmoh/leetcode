package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPathCrossing(t *testing.T) {
	t.Run("should return false", func(t *testing.T) {
		path := "NES"
		assert.False(t, isPathCrossing(path))
	})

	t.Run("should return true", func(t *testing.T) {
		path := "NESWW"
		assert.True(t, isPathCrossing(path))
	})
}
