package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividePlayers(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		skill := []int{3, 2, 5, 1, 3, 4}
		expected := int64(22)

		assert.Equal(t, expected, dividePlayers(skill))
	})

	t.Run("test case 2", func(t *testing.T) {
		skill := []int{3, 4}
		expected := int64(12)

		assert.Equal(t, expected, dividePlayers(skill))
	})

	t.Run("test case 1", func(t *testing.T) {
		skill := []int{1, 1, 2, 3}
		expected := int64(-1)

		assert.Equal(t, expected, dividePlayers(skill))
	})
}
