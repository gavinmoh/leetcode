package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumLength(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "ca"
		expected := 2

		assert.Equal(t, expected, minimumLength(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "cabaabac"
		expected := 0

		assert.Equal(t, expected, minimumLength(s))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "aabccabba"
		expected := 3

		assert.Equal(t, expected, minimumLength(s))
	})

	t.Run("test case 4", func(t *testing.T) {
		s := "abbbbbbbbbbbbbbbbbbba"
		expected := 0

		assert.Equal(t, expected, minimumLength(s))
	})
}
