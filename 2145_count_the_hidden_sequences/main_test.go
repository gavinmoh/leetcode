package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfArrays(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		differences := []int{1, -3, 4}
		lower := 1
		upper := 6
		expected := 2

		assert.Equal(t, expected, numberOfArrays(differences, lower, upper))
	})

	t.Run("test case 2", func(t *testing.T) {
		differences := []int{3, -4, 5, 1, -2}
		lower := -4
		upper := 5
		expected := 4

		assert.Equal(t, expected, numberOfArrays(differences, lower, upper))
	})

	t.Run("test case 3", func(t *testing.T) {
		differences := []int{4, -7, 2}
		lower := 3
		upper := 6
		expected := 0

		assert.Equal(t, expected, numberOfArrays(differences, lower, upper))
	})
}
