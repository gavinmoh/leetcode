package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinAddToMakeValid(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "())"
		expected := 1

		assert.Equal(t, expected, minAddToMakeValid(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "((("
		expected := 3

		assert.Equal(t, expected, minAddToMakeValid(s))
	})
}
