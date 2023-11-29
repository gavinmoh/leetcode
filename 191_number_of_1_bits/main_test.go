package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHammingWeight(t *testing.T) {
	t.Run("should return 3", func(t *testing.T) {
		num := uint32(11)
		expected := 3
		assert.Equal(t, expected, hammingWeight(num))
	})

	t.Run("should return 1", func(t *testing.T) {
		num := uint32(128)
		expected := 1
		assert.Equal(t, expected, hammingWeight(num))
	})

	t.Run("should return 31", func(t *testing.T) {
		num := uint32(4294967293)
		expected := 31
		assert.Equal(t, expected, hammingWeight(num))
	})
}
