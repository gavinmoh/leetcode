package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructDistancedSequence(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 3
		expected := []int{3, 1, 2, 3, 2}

		assert.Equal(t, expected, constructDistancedSequence(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 5
		expected := []int{5, 3, 1, 4, 3, 5, 2, 4, 2}

		assert.Equal(t, expected, constructDistancedSequence(n))
	})
}
