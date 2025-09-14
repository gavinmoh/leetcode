package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpellCheckec(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		wordlist := []string{"KiTe", "kite", "hare", "Hare"}
		queries := []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}
		expected := []string{"kite", "KiTe", "KiTe", "Hare", "hare", "", "", "KiTe", "", "KiTe"}

		assert.Equal(t, expected, spellchecker(wordlist, queries))
	})

	t.Run("test case 2", func(t *testing.T) {
		wordlist := []string{"yellow"}
		queries := []string{"YellOw"}
		expected := []string{"yellow"}

		assert.Equal(t, expected, spellchecker(wordlist, queries))
	})

	t.Run("test case 3", func(t *testing.T) {
		wordlist := []string{"ae", "aa"}
		queries := []string{"UU"}
		expected := []string{"ae"}

		assert.Equal(t, expected, spellchecker(wordlist, queries))
	})
}
