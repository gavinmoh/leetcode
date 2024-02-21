package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeBitwiseAnd(t *testing.T) {
	t.Run("should return 4", func(t *testing.T) {
		left := 5
		right := 7
		expected := 4

		assert.Equal(t, expected, rangeBitwiseAnd(left, right))
	})

	t.Run("should return 0", func(t *testing.T) {
		left := 0
		right := 0
		expected := 0

		assert.Equal(t, expected, rangeBitwiseAnd(left, right))
	})

	t.Run("should return 0", func(t *testing.T) {
		left := 1
		right := 2147483647
		expected := 0

		assert.Equal(t, expected, rangeBitwiseAnd(left, right))
	})
}
