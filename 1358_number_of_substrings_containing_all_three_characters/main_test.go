package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfSubstrings(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "abcabc"
		expected := 10

		assert.Equal(t, expected, numberOfSubstrings(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "aaacb"
		expected := 3

		assert.Equal(t, expected, numberOfSubstrings(s))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "abc"
		expected := 1

		assert.Equal(t, expected, numberOfSubstrings(s))
	})
}
