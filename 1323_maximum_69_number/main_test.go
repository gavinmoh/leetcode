package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximum69Number(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		num := 9669
		expected := 9969

		assert.Equal(t, expected, maximum69Number(num))
	})

	t.Run("test case 2", func(t *testing.T) {
		num := 9996
		expected := 9999

		assert.Equal(t, expected, maximum69Number(num))
	})

	t.Run("test case 3", func(t *testing.T) {
		num := 9999
		expected := 9999

		assert.Equal(t, expected, maximum69Number(num))
	})
}
