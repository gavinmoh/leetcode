package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountOperations(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		num1 := 2
		num2 := 3
		expected := 3

		assert.Equal(t, expected, countOperations(num1, num2))
	})

	t.Run("test case 2", func(t *testing.T) {
		num1 := 10
		num2 := 10
		expected := 1

		assert.Equal(t, expected, countOperations(num1, num2))
	})
}
