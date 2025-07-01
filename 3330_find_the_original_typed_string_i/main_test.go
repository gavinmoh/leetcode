package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPossibleStringCount(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		word := "abbcccc"
		expected := 5

		assert.Equal(t, expected, possibleStringCount(word))
	})

	t.Run("test case 2", func(t *testing.T) {
		word := "abcd"
		expected := 1

		assert.Equal(t, expected, possibleStringCount(word))
	})

	t.Run("test case 3", func(t *testing.T) {
		word := "aaaa"
		expected := 4

		assert.Equal(t, expected, possibleStringCount(word))
	})
}
