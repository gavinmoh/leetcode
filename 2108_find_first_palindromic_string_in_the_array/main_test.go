package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstPalindrome(t *testing.T) {
	t.Run("should return ada", func(t *testing.T) {
		words := []string{"abc", "car", "ada", "racecar", "cool"}
		expected := "ada"

		assert.Equal(t, expected, firstPalindrome(words))
	})

	t.Run("should return racecar", func(t *testing.T) {
		words := []string{"notapalindrome", "racecar"}
		expected := "racecar"

		assert.Equal(t, expected, firstPalindrome(words))
	})

	t.Run("should return empty string", func(t *testing.T) {
		words := []string{"def", "ghi"}
		expected := ""

		assert.Equal(t, expected, firstPalindrome(words))
	})
}
