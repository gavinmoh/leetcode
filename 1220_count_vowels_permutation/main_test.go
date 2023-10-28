package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountVowelPermutation(t *testing.T) {
	t.Run("should return 5", func(t *testing.T) {
		n := 1
		expected := 5
		assert.Equal(t, expected, countVowelPermutation(n))
	})

	t.Run("should return 10", func(t *testing.T) {
		n := 2
		expected := 10
		assert.Equal(t, expected, countVowelPermutation(n))
	})

	t.Run("should return 68", func(t *testing.T) {
		n := 5
		expected := 68
		assert.Equal(t, expected, countVowelPermutation(n))
	})

}
