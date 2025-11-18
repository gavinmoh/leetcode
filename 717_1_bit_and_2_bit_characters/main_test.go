package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsOneBitCharacter(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		bits := []int{1, 0, 0}

		assert.True(t, isOneBitCharacter(bits))
	})

	t.Run("test case 2", func(t *testing.T) {
		bits := []int{1, 1, 1, 0}

		assert.False(t, isOneBitCharacter(bits))
	})
}
