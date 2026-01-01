package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		digits := []int{1, 2, 3}
		expected := []int{1, 2, 4}

		assert.Equal(t, expected, plusOne(digits))
	})

	t.Run("test case 2", func(t *testing.T) {
		digits := []int{4, 3, 2, 1}
		expected := []int{4, 3, 2, 2}

		assert.Equal(t, expected, plusOne(digits))
	})

	t.Run("test case 3", func(t *testing.T) {
		digits := []int{9}
		expected := []int{1, 0}

		assert.Equal(t, expected, plusOne(digits))
	})
}
