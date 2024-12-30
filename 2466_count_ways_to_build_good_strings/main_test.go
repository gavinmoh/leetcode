package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountGoodStrings(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		low := 3
		high := 3
		zero := 1
		one := 1
		expected := 8

		assert.Equal(t, expected, countGoodStrings(low, high, zero, one))
	})

	t.Run("test case 2", func(t *testing.T) {
		low := 2
		high := 3
		zero := 1
		one := 2
		expected := 5

		assert.Equal(t, expected, countGoodStrings(low, high, zero, one))
	})
}
