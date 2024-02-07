package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrequencySort(t *testing.T) {
	t.Run("should return eert", func(t *testing.T) {
		s := "tree"
		expected := []string{"eert", "eetr"}

		assert.Contains(t, expected, frequencySort(s))
	})

	t.Run("should return aaaccc", func(t *testing.T) {
		s := "cccaaa"
		expected := []string{"aaaccc", "cccaaa"}

		assert.Contains(t, expected, frequencySort(s))
	})

	t.Run("should return bbAa", func(t *testing.T) {
		s := "Aabb"
		expected := []string{"bbAa", "bbaA"}

		assert.Contains(t, expected, frequencySort(s))
	})
}
