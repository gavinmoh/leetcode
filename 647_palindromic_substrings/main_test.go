package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSubstrings(t *testing.T) {
	t.Run("should return 3", func(t *testing.T) {
		s := "abc"
		expected := 3

		assert.Equal(t, expected, countSubstrings(s))
	})

	t.Run("should return 6", func(t *testing.T) {
		s := "aaa"
		expected := 6

		assert.Equal(t, expected, countSubstrings(s))
	})
}
