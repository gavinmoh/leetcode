package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingRolls(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		rolls := []int{3, 2, 4, 3}
		mean := 4
		n := 2
		expected := []int{6, 6}

		assert.Equal(t, expected, missingRolls(rolls, mean, n))
	})

	t.Run("test case 2", func(t *testing.T) {
		rolls := []int{1, 5, 6}
		mean := 3
		n := 4
		expected := []int{3, 2, 2, 2}

		assert.Equal(t, expected, missingRolls(rolls, mean, n))
	})

	t.Run("test case 3", func(t *testing.T) {
		rolls := []int{1, 2, 3, 4}
		mean := 6
		n := 4
		expected := []int{}

		assert.Equal(t, expected, missingRolls(rolls, mean, n))
	})

	t.Run("test case 4", func(t *testing.T) {
		rolls := []int{4, 5, 6, 2, 3, 6, 5, 4, 6, 4, 5, 1, 6, 3, 1, 4, 5, 5, 3, 2, 3, 5, 3, 2, 1, 5, 4, 3, 5, 1, 5}
		mean := 4
		n := 40
		expected := []int{5, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}

		assert.Equal(t, expected, missingRolls(rolls, mean, n))
	})
}
