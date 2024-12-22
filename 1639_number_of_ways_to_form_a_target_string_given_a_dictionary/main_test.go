package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumWays(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		words := []string{"acca", "bbbb", "caca"}
		target := "aba"
		expected := 6

		assert.Equal(t, expected, numWays(words, target))
	})

	t.Run("test case 2", func(t *testing.T) {
		words := []string{"abba", "baab"}
		target := "bab"
		expected := 4

		assert.Equal(t, expected, numWays(words, target))
	})
}
