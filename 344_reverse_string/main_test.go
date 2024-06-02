package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := []byte{'h', 'e', 'l', 'l', 'o'}
		expected := []byte{'o', 'l', 'l', 'e', 'h'}
		reverseString(s)

		assert.Equal(t, expected, s)
	})

	t.Run("test case 2", func(t *testing.T) {
		s := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
		expected := []byte{'h', 'a', 'n', 'n', 'a', 'H'}
		reverseString(s)

		assert.Equal(t, expected, s)
	})
}
