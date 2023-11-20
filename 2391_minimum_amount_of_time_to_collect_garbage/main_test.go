package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGarbageCollection(t *testing.T) {
	t.Run("should return 21", func(t *testing.T) {
		garbage := []string{"G", "P", "GP", "GG"}
		travel := []int{2, 4, 3}
		expected := 21
		assert.Equal(t, expected, garbageCollection(garbage, travel))
	})

	t.Run("should return 37", func(t *testing.T) {
		garbage := []string{"MMM", "PGM", "GP"}
		travel := []int{3, 10}
		expected := 37
		assert.Equal(t, expected, garbageCollection(garbage, travel))
	})
}
