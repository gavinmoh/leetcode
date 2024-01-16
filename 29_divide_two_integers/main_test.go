package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	t.Run("should return 3", func(t *testing.T) {
		dividend := 10
		divisor := 3
		expected := 3
		assert.Equal(t, expected, divide(dividend, divisor))
	})

	t.Run("should return -2", func(t *testing.T) {
		dividend := 7
		divisor := -3
		expected := -2
		assert.Equal(t, expected, divide(dividend, divisor))
	})

	t.Run("should return 2147483647", func(t *testing.T) {
		dividend := -2147483648
		divisor := -1
		expected := 2147483647
		assert.Equal(t, expected, divide(dividend, divisor))
	})

	t.Run("should return -2147483648", func(t *testing.T) {
		dividend := -2147483648
		divisor := 1
		expected := -2147483648
		assert.Equal(t, expected, divide(dividend, divisor))
	})

}
