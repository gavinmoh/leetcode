package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumEnergy(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		energy := []int{5, 2, -10, -5, 1}
		k := 3
		expected := 3

		assert.Equal(t, expected, maximumEnergy(energy, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		energy := []int{-2, -3, -1}
		k := 2
		expected := -1

		assert.Equal(t, expected, maximumEnergy(energy, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		energy := []int{8, -5}
		k := 1
		expected := 3

		assert.Equal(t, expected, maximumEnergy(energy, k))
	})
}
