package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrangePrinter(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "aaabbb"
		expected := 2

		assert.Equal(t, expected, strangePrinter(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "aba"
		expected := 2

		assert.Equal(t, expected, strangePrinter(s))
	})
}
