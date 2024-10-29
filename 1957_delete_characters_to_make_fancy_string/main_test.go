package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeFancyString(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "leeetcode"
		expected := "leetcode"

		assert.Equal(t, expected, makeFancyString(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "aaabaaaa"
		expected := "aabaa"

		assert.Equal(t, expected, makeFancyString(s))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "aab"
		expected := "aab"

		assert.Equal(t, expected, makeFancyString(s))
	})
}
