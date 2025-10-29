package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestNumber(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 5
		expected := 7

		assert.Equal(t, expected, smallestNumber(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 10
		expected := 15

		assert.Equal(t, expected, smallestNumber(n))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 3
		expected := 3

		assert.Equal(t, expected, smallestNumber(n))
	})
}
