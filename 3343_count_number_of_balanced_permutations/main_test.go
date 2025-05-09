package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBalancedPermutations(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		num := "123"
		expected := 2

		assert.Equal(t, expected, countBalancedPermutations(num))
	})

	t.Run("test case 2", func(t *testing.T) {
		num := "112"
		expected := 1

		assert.Equal(t, expected, countBalancedPermutations(num))
	})

	t.Run("test case 3", func(t *testing.T) {
		num := "12345"
		expected := 0

		assert.Equal(t, expected, countBalancedPermutations(num))
	})
}
