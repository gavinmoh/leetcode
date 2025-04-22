package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdealArrays(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 2
		maxValue := 5
		expected := 10

		assert.Equal(t, expected, idealArrays(n, maxValue))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 5
		maxValue := 3
		expected := 11

		assert.Equal(t, expected, idealArrays(n, maxValue))
	})
}
