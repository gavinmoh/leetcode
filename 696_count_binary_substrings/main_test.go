package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBinarySubstrings(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "00110011"
		expected := 6

		assert.Equal(t, expected, countBinarySubstrings(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "10101"
		expected := 4

		assert.Equal(t, expected, countBinarySubstrings(s))
	})
}
