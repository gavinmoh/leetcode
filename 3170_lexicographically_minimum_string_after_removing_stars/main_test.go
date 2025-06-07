package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClearStars(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "aaba*"
		expected := "aab"

		assert.Equal(t, expected, clearStars(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "abc"
		expected := "abc"

		assert.Equal(t, expected, clearStars(s))
	})
}
