package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSteps(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "1101"
		expected := 6

		assert.Equal(t, expected, numSteps(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "10"
		expected := 1

		assert.Equal(t, expected, numSteps(s))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "1"
		expected := 0

		assert.Equal(t, expected, numSteps(s))
	})
}
