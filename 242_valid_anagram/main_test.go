package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	// t.Run("should return true", func(t *testing.T) {
	// 	s1 := "anagram"
	// 	s2 := "nagaram"
	// 	assert.True(t, isAnagram(s1, s2))
	// })

	t.Run("should return false", func(t *testing.T) {
		s1 := "rat"
		s2 := "car"
		assert.False(t, isAnagram(s1, s2))
	})
}
