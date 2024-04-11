package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveKDigits(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		num := "1432219"
		k := 3
		expected := "1219"

		assert.Equal(t, expected, removeKdigits(num, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		num := "10200"
		k := 1
		expected := "200"

		assert.Equal(t, expected, removeKdigits(num, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		num := "10"
		k := 2
		expected := "0"

		assert.Equal(t, expected, removeKdigits(num, k))
	})

	t.Run("test case 4", func(t *testing.T) {
		num := "9"
		k := 1
		expected := "0"

		assert.Equal(t, expected, removeKdigits(num, k))
	})

	t.Run("test case 5", func(t *testing.T) {
		num := "1234567890"
		k := 9
		expected := "0"

		assert.Equal(t, expected, removeKdigits(num, k))
	})
}
