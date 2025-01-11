package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanConstruct(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "annabelle"
		k := 2

		assert.True(t, canConstruct(s, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "leetcode"
		k := 3

		assert.False(t, canConstruct(s, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "true"
		k := 4

		assert.True(t, canConstruct(s, k))
	})
}
