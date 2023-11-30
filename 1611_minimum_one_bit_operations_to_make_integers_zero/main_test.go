package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumOneBitOperations(t *testing.T) {
	t.Run("should return 2", func(t *testing.T) {
		n := 3
		expected := 2
		assert.Equal(t, expected, minimumOneBitOperations(n))
	})

	t.Run("should return 4", func(t *testing.T) {
		n := 6
		expected := 4
		assert.Equal(t, expected, minimumOneBitOperations(n))
	})
}
