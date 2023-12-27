package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCost(t *testing.T) {
	t.Run("should return 3", func(t *testing.T) {
		colors := "abaac"
		neededTime := []int{1, 2, 3, 4, 5}
		expected := 3
		assert.Equal(t, expected, minCost(colors, neededTime))
	})

	t.Run("should return 0", func(t *testing.T) {
		colors := "abc"
		neededTime := []int{1, 2, 3}
		expected := 0
		assert.Equal(t, expected, minCost(colors, neededTime))
	})

	t.Run("should return 2", func(t *testing.T) {
		colors := "aabaa"
		neededTime := []int{1, 2, 3, 4, 1}
		expected := 2
		assert.Equal(t, expected, minCost(colors, neededTime))
	})
}
