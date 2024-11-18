package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTakeCharacters(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "aabaaaacaabc"
		k := 2
		expected := 8

		assert.Equal(t, expected, takeCharacters(s, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "a"
		k := 1
		expected := -1

		assert.Equal(t, expected, takeCharacters(s, k))
	})
}
