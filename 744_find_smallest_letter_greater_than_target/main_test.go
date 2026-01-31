package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextGreatestLetter(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		letters := []byte{'c', 'f', 'j'}
		target := byte('a')
		expected := byte('c')

		assert.Equal(t, expected, nextGreatestLetter(letters, target))
	})

	t.Run("test case 2", func(t *testing.T) {
		letters := []byte{'c', 'f', 'j'}
		target := byte('c')
		expected := byte('f')

		assert.Equal(t, expected, nextGreatestLetter(letters, target))
	})

	t.Run("test case 3", func(t *testing.T) {
		letters := []byte{'x', 'x', 'y', 'y'}
		target := byte('z')
		expected := byte('x')

		assert.Equal(t, expected, nextGreatestLetter(letters, target))
	})
}
