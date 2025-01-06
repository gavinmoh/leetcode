package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringMatching(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		words := []string{"mass", "as", "hero", "superhero"}
		expected := []string{"as", "hero"}

		assert.Equal(t, expected, stringMatching(words))
	})

	t.Run("test case 2", func(t *testing.T) {
		words := []string{"leetcode", "et", "code"}
		expected := []string{"et", "code"}

		assert.Equal(t, expected, stringMatching(words))
	})

	t.Run("test case 3", func(t *testing.T) {
		words := []string{"blue", "green", "bu"}
		expected := []string{}

		assert.Equal(t, expected, stringMatching(words))
	})

	t.Run("test case 4", func(t *testing.T) {
		words := []string{"leetcoder", "leetcode", "od", "hamlet", "am"}
		expected := []string{"leetcode", "od", "am"}

		assert.Equal(t, expected, stringMatching(words))
	})
}
