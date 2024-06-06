package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNStraightHand(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		hand := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
		groupSize := 3

		assert.True(t, isNStraightHand(hand, groupSize))
	})

	t.Run("test case 2", func(t *testing.T) {
		hand := []int{1, 2, 3, 4, 5}
		groupSize := 4

		assert.False(t, isNStraightHand(hand, groupSize))
	})

	t.Run("test case 3", func(t *testing.T) {
		hand := []int{6, 6, 6, 7, 5, 5, 6, 5}
		groupSize := 8

		assert.False(t, isNStraightHand(hand, groupSize))
	})
}
