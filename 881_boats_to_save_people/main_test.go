package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumRescueBoats(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		people := []int{1, 2}
		limit := 3
		expected := 1

		assert.Equal(t, expected, numRescueBoats(people, limit))
	})

	t.Run("test case 2", func(t *testing.T) {
		people := []int{3, 2, 2, 1}
		limit := 3
		expected := 3

		assert.Equal(t, expected, numRescueBoats(people, limit))
	})

	t.Run("test case 3", func(t *testing.T) {
		people := []int{3, 5, 3, 4}
		limit := 5
		expected := 4

		assert.Equal(t, expected, numRescueBoats(people, limit))
	})

	t.Run("test case 4", func(t *testing.T) {
		people := []int{2, 49, 10, 7, 11, 41, 47, 2, 22, 6, 13, 12, 33, 18, 10, 26, 2, 6, 50, 10}
		limit := 50
		expected := 11

		assert.Equal(t, expected, numRescueBoats(people, limit))
	})
}
