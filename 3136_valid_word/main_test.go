package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		word := "234Adas"

		assert.True(t, isValid(word))
	})

	t.Run("test case 2", func(t *testing.T) {
		word := "b3"

		assert.False(t, isValid(word))
	})

	t.Run("test case 3", func(t *testing.T) {
		word := "a3$e"

		assert.False(t, isValid(word))
	})
}
