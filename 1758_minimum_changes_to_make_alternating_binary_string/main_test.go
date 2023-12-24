package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinOperations(t *testing.T) {
	t.Run("should return 1", func(t *testing.T) {
		s := "0100"
		expected := 1
		assert.Equal(t, expected, minOperations(s))
	})

	t.Run("should return 0", func(t *testing.T) {
		s := "10"
		expected := 0
		assert.Equal(t, expected, minOperations(s))
	})

	t.Run("should return 2", func(t *testing.T) {
		s := "1111"
		expected := 2
		assert.Equal(t, expected, minOperations(s))
	})
}
