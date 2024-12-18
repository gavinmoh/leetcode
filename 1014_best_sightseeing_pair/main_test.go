package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxScoreSightseeingPair(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		values := []int{8, 1, 5, 2, 6}
		expected := 11

		assert.Equal(t, expected, maxScoreSightseeingPair(values))
	})

	t.Run("test case 2", func(t *testing.T) {
		values := []int{1, 2}
		expected := 2

		assert.Equal(t, expected, maxScoreSightseeingPair(values))
	})

	t.Run("test case 3", func(t *testing.T) {
		values := []int{1, 3, 5}
		expected := 7

		assert.Equal(t, expected, maxScoreSightseeingPair(values))
	})
}
