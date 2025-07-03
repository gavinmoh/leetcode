package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthCharacter(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		k := 5
		expected := byte('b')

		assert.Equal(t, expected, kthCharacter(k))
	})

	t.Run("test case 2", func(t *testing.T) {
		k := 10
		expected := byte('c')

		assert.Equal(t, expected, kthCharacter(k))
	})
}
