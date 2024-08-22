package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitwiseComplement(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 5
		expected := 2

		assert.Equal(t, expected, bitwiseComplement(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 7
		expected := 0

		assert.Equal(t, expected, bitwiseComplement(n))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 10
		expected := 5

		assert.Equal(t, expected, bitwiseComplement(n))
	})

	t.Run("test case 4", func(t *testing.T) {
		n := 0
		expected := 1

		assert.Equal(t, expected, bitwiseComplement(n))
	})
}
