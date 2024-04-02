package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIsomorphic(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s1 := "egg"
		s2 := "add"

		assert.True(t, isIsomorphic(s1, s2))
	})

	t.Run("test case 2", func(t *testing.T) {
		s1 := "foo"
		s2 := "bar"

		assert.False(t, isIsomorphic(s1, s2))
	})

	t.Run("test case 3", func(t *testing.T) {
		s1 := "paper"
		s2 := "title"

		assert.True(t, isIsomorphic(s1, s2))
	})
}
