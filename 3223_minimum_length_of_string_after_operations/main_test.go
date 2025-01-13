package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumLength(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "abaacbcbb"
		expected := 5

		assert.Equal(t, expected, minimumLength(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "aa"
		expected := 2

		assert.Equal(t, expected, minimumLength(s))
	})
}
