package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumInvitations(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		favorite := []int{2, 2, 1, 2}
		expected := 3

		assert.Equal(t, expected, maximumInvitations(favorite))
	})

	t.Run("test case 2", func(t *testing.T) {
		favorite := []int{1, 2, 0}
		expected := 3

		assert.Equal(t, expected, maximumInvitations(favorite))
	})

	t.Run("test case 3", func(t *testing.T) {
		favorite := []int{3, 0, 1, 4, 1}
		expected := 4

		assert.Equal(t, expected, maximumInvitations(favorite))
	})
}
