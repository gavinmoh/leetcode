package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "(1+(2*3)+((8)/4))+1"
		expected := 3

		assert.Equal(t, expected, maxDepth(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "(1)+((2))+(((3)))"
		expected := 3

		assert.Equal(t, expected, maxDepth(s))
	})
}
