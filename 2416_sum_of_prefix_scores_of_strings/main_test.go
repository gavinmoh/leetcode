package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumPrefixScores(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		words := []string{"abc", "ab", "bc", "b"}
		expected := []int{5, 4, 3, 2}

		assert.Equal(t, expected, sumPrefixScores(words))
	})

	t.Run("test case 2", func(t *testing.T) {
		words := []string{"abcd"}
		expected := []int{4}

		assert.Equal(t, expected, sumPrefixScores(words))
	})

}
