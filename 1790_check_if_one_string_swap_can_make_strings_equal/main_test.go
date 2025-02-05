package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreAlmostEqual(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s1 := "bank"
		s2 := "kanb"

		assert.True(t, areAlmostEqual(s1, s2))
	})

	t.Run("test case 2", func(t *testing.T) {
		s1 := "attack"
		s2 := "defend"

		assert.False(t, areAlmostEqual(s1, s2))
	})

	t.Run("test case 3", func(t *testing.T) {
		s1 := "kelb"
		s2 := "kelb"

		assert.True(t, areAlmostEqual(s1, s2))
	})

	t.Run("test case 4", func(t *testing.T) {
		s1 := "abcd"
		s2 := "dcba"

		assert.False(t, areAlmostEqual(s1, s2))
	})
}
