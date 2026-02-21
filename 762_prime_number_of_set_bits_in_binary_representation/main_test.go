package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPrimeSetBits(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		left := 6
		right := 10
		expected := 4

		assert.Equal(t, expected, countPrimeSetBits(left, right))
	})

	t.Run("test case 2", func(t *testing.T) {
		left := 10
		right := 15
		expected := 5

		assert.Equal(t, expected, countPrimeSetBits(left, right))
	})
}
