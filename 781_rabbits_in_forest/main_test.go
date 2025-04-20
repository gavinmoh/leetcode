package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumRabbits(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		answers := []int{1, 1, 2}
		expected := 5

		assert.Equal(t, expected, numRabbits(answers))
	})

	t.Run("test case 2", func(t *testing.T) {
		answers := []int{10, 10, 10}
		expected := 11

		assert.Equal(t, expected, numRabbits(answers))
	})
}
