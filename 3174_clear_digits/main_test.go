package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClearDigits(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "abc"
		expected := "abc"

		assert.Equal(t, expected, clearDigits(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "cb34"
		expected := ""

		assert.Equal(t, expected, clearDigits(s))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "b4y6"
		expected := ""

		assert.Equal(t, expected, clearDigits(s))
	})
}
