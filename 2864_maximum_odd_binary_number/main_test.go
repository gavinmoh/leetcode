package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumOddBinaryNumber(t *testing.T) {
	t.Run("should return 001", func(t *testing.T) {
		s := "010"
		expected := "001"

		assert.Equal(t, expected, maximumOddBinaryNumber(s))
	})

	t.Run("should return 1001", func(t *testing.T) {
		s := "0101"
		expected := "1001"

		assert.Equal(t, expected, maximumOddBinaryNumber(s))
	})
}
