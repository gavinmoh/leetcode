package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCollisions(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		directions := "RLRSLL"
		expected := 5

		assert.Equal(t, expected, countCollisions(directions))
	})

	t.Run("test case 2", func(t *testing.T) {
		directions := "LLRR"
		expected := 0

		assert.Equal(t, expected, countCollisions(directions))
	})
}
