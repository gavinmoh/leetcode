package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcatenatedBinary(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 1
		expected := 1

		assert.Equal(t, expected, concatenatedBinary(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 3
		expected := 27

		assert.Equal(t, expected, concatenatedBinary(n))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 12
		expected := 505379714

		assert.Equal(t, expected, concatenatedBinary(n))
	})

	t.Run("test case 4", func(t *testing.T) {
		n := 42
		expected := 727837408

		assert.Equal(t, expected, concatenatedBinary(n))
	})
}
