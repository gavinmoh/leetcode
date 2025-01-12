package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimizeXor(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		num1 := 3
		num2 := 5
		expected := 3

		assert.Equal(t, expected, minimizeXor(num1, num2))
	})

	t.Run("test case 2", func(t *testing.T) {
		num1 := 1
		num2 := 12
		expected := 3

		assert.Equal(t, expected, minimizeXor(num1, num2))
	})
}
