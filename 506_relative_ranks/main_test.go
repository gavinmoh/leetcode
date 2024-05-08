package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRelativeRanks(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		score := []int{5, 4, 3, 2, 1}
		expected := []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}

		assert.Equal(t, expected, findRelativeRanks(score))
	})

	t.Run("test case 2", func(t *testing.T) {
		score := []int{10, 3, 8, 9, 4}
		expected := []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"}

		assert.Equal(t, expected, findRelativeRanks(score))
	})
}
