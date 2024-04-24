package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTribonacci(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 4
		expected := 4

		assert.Equal(t, expected, tribonacci(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 25
		expected := 1_389_537

		assert.Equal(t, expected, tribonacci(n))
	})
}
