package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDifference(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "aaaaabbc"
		expected := 3

		assert.Equal(t, expected, maxDifference(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "abcabcab"
		expected := 1

		assert.Equal(t, expected, maxDifference(s))
	})
}
