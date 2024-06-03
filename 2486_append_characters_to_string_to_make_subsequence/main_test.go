package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendCharacters(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s1 := "coaching"
		s2 := "coding"
		expected := 4

		assert.Equal(t, expected, appendCharacters(s1, s2))
	})

	t.Run("test case 2", func(t *testing.T) {
		s1 := "abcde"
		s2 := "a"
		expected := 0

		assert.Equal(t, expected, appendCharacters(s1, s2))
	})

	t.Run("test case 3", func(t *testing.T) {
		s1 := "z"
		s2 := "abcde"
		expected := 5

		assert.Equal(t, expected, appendCharacters(s1, s2))
	})

	t.Run("test case 4", func(t *testing.T) {
		s1 := "vrykt"
		s2 := "rkge"
		expected := 2

		assert.Equal(t, expected, appendCharacters(s1, s2))
	})
}
