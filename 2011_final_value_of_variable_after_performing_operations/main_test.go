package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFinalValueAfterOperations(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		operations := []string{"--X", "X++", "X++"}
		expected := 1

		assert.Equal(t, expected, finalValueAfterOperations(operations))
	})

	t.Run("test case 2", func(t *testing.T) {
		operations := []string{"++X", "++X", "X++"}
		expected := 3

		assert.Equal(t, expected, finalValueAfterOperations(operations))
	})

	t.Run("test case 3", func(t *testing.T) {
		operations := []string{"X++", "++X", "--X", "X--"}
		expected := 0

		assert.Equal(t, expected, finalValueAfterOperations(operations))
	})
}
