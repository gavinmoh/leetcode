package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveOccurrences(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "daabcbaabcbc"
		part := "abc"
		expected := "dab"

		assert.Equal(t, expected, removeOccurrences(s, part))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "axxxxyyyyb"
		part := "xy"
		expected := "ab"

		assert.Equal(t, expected, removeOccurrences(s, part))
	})
}
