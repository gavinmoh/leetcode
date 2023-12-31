package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxLengthBetweenEqualCharacters(t *testing.T) {
	t.Run("should return 0", func(t *testing.T) {
		s := "aa"
		expected := 0
		assert.Equal(t, expected, maxLengthBetweenEqualCharacters(s))
	})

	t.Run("should return 2", func(t *testing.T) {
		s := "abca"
		expected := 2
		assert.Equal(t, expected, maxLengthBetweenEqualCharacters(s))
	})

	t.Run("should return -1", func(t *testing.T) {
		s := "cbzxy"
		expected := -1
		assert.Equal(t, expected, maxLengthBetweenEqualCharacters(s))
	})
}
