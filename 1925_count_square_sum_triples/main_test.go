package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountTriples(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 5
		expected := 2

		assert.Equal(t, expected, countTriples(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 10
		expected := 4

		assert.Equal(t, expected, countTriples(n))
	})
}
