package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumZero(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 5
		expected := []int{-1, 1, -3, 3, 0}

		assert.Equal(t, expected, sumZero(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 3
		expected := []int{-1, 1, 0}

		assert.Equal(t, expected, sumZero(n))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 1
		expected := []int{0}

		assert.Equal(t, expected, sumZero(n))
	})
}
