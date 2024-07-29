package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumTeams(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		rating := []int{2, 5, 3, 4, 1}
		expected := 3

		assert.Equal(t, expected, numTeams(rating))
	})

	t.Run("test case 2", func(t *testing.T) {
		rating := []int{2, 1, 3}
		expected := 0

		assert.Equal(t, expected, numTeams(rating))
	})

	t.Run("test case 3", func(t *testing.T) {
		rating := []int{1, 2, 3, 4}
		expected := 4

		assert.Equal(t, expected, numTeams(rating))
	})
}
