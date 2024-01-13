package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSteps(t *testing.T) {
	t.Run("should return 1", func(t *testing.T) {
		s1 := "bab"
		s2 := "aba"
		expected := 1
		assert.Equal(t, expected, minSteps(s1, s2))
	})

	t.Run("should return 5", func(t *testing.T) {
		s1 := "leetcode"
		s2 := "practice"
		expected := 5
		assert.Equal(t, expected, minSteps(s1, s2))
	})

	t.Run("should return 0", func(t *testing.T) {
		s1 := "anagram"
		s2 := "mangaar"
		expected := 0
		assert.Equal(t, expected, minSteps(s1, s2))
	})
}
