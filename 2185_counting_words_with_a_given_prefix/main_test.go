package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixCount(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		words := []string{"pay", "attention", "practice", "attend"}
		pref := "at"
		expected := 2

		assert.Equal(t, expected, prefixCount(words, pref))
	})

	t.Run("test case 2", func(t *testing.T) {
		words := []string{"leetcode", "win", "loops", "success"}
		pref := "code"
		expected := 0

		assert.Equal(t, expected, prefixCount(words, pref))
	})
}
