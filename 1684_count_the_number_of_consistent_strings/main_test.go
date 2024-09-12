package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountConsistentStrings(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		allowed := "ab"
		words := []string{"ad", "bd", "aaab", "baa", "badab"}
		expected := 2

		assert.Equal(t, expected, countConsistentStrings(allowed, words))
	})

	t.Run("test case 2", func(t *testing.T) {
		allowed := "abc"
		words := []string{"a", "b", "c", "ab", "ac", "bc", "abc"}
		expected := 7

		assert.Equal(t, expected, countConsistentStrings(allowed, words))
	})

	t.Run("test case 3", func(t *testing.T) {
		allowed := "cad"
		words := []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}
		expected := 4

		assert.Equal(t, expected, countConsistentStrings(allowed, words))
	})
}
