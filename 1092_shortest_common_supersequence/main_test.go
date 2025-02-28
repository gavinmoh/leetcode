package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestCommonSupersequence(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		str1 := "abac"
		str2 := "cab"
		expected := "cabac"

		assert.Equal(t, expected, shortestCommonSupersequence(str1, str2))
	})

	t.Run("test case 2", func(t *testing.T) {
		str1 := "aaaaaaaa"
		str2 := "aaaaaaaa"
		expected := "aaaaaaaa"

		assert.Equal(t, expected, shortestCommonSupersequence(str1, str2))
	})

	t.Run("test case 3", func(t *testing.T) {
		str1 := "bbbaaaba"
		str2 := "bbababbb"
		expected := "bbbaaababbb"

		assert.Equal(t, expected, shortestCommonSupersequence(str1, str2))
	})
}
