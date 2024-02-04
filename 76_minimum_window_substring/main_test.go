package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinWindow(t *testing.T) {
	t.Run("should return BANC", func(t *testing.T) {
		s1 := "ADOBECODEBANC"
		s2 := "ABC"
		expected := "BANC"

		assert.Equal(t, expected, minWindow(s1, s2))
	})

	t.Run("should return a", func(t *testing.T) {
		s1 := "a"
		s2 := "a"
		expected := "a"

		assert.Equal(t, expected, minWindow(s1, s2))
	})

	t.Run("should return empty string", func(t *testing.T) {
		s1 := "a"
		s2 := "aa"
		expected := ""

		assert.Equal(t, expected, minWindow(s1, s2))
	})
}
