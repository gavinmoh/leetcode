package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessfulPairs(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		spells := []int{5, 1, 3}
		potions := []int{1, 2, 3, 4, 5}
		success := int64(7)
		expected := []int{4, 0, 3}

		assert.Equal(t, expected, successfulPairs(spells, potions, success))
	})

	t.Run("test case 2", func(t *testing.T) {
		spells := []int{3, 1, 2}
		potions := []int{8, 5, 8}
		success := int64(16)
		expected := []int{2, 0, 2}

		assert.Equal(t, expected, successfulPairs(spells, potions, success))
	})
}
