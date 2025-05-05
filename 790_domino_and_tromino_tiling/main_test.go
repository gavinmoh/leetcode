package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumTilings(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 3
		expected := 5

		assert.Equal(t, expected, numTilings(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 1
		expected := 1

		assert.Equal(t, expected, numTilings(n))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 60
		expected := 882347204

		assert.Equal(t, expected, numTilings(n))
	})
}
