package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryGap(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 22
		expected := 2

		assert.Equal(t, expected, binaryGap(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 8
		expected := 0

		assert.Equal(t, expected, binaryGap(n))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 5
		expected := 2

		assert.Equal(t, expected, binaryGap(n))
	})
}
