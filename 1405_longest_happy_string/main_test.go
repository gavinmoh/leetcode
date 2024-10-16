package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestDiverseString(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		a := 1
		b := 1
		c := 7
		expected := "ccaccbcc"

		assert.Equal(t, expected, longestDiverseString(a, b, c))
	})

	t.Run("test case 2", func(t *testing.T) {
		a := 7
		b := 1
		c := 0
		expected := "aabaa"

		assert.Equal(t, expected, longestDiverseString(a, b, c))
	})

	t.Run("test case 3", func(t *testing.T) {
		a := 1
		b := 4
		c := 5
		expected := "ccbbccbacb"

		assert.Equal(t, expected, longestDiverseString(a, b, c))
	})

	t.Run("test case 4", func(t *testing.T) {
		a := 0
		b := 0
		c := 7
		expected := "cc"

		assert.Equal(t, expected, longestDiverseString(a, b, c))
	})
}
