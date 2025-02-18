package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestNumber(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		pattern := "IIIDIDDD"
		expected := "123549876"

		assert.Equal(t, expected, smallestNumber(pattern))
	})

	t.Run("test case 2", func(t *testing.T) {
		pattern := "DDD"
		expected := "4321"

		assert.Equal(t, expected, smallestNumber(pattern))
	})
}
