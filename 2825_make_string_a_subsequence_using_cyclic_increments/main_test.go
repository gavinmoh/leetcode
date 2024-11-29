package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanMakeSubsequence(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		str1 := "abc"
		str2 := "ad"

		assert.True(t, canMakeSubsequence(str1, str2))
	})

	t.Run("test case 2", func(t *testing.T) {
		str1 := "zc"
		str2 := "ad"

		assert.True(t, canMakeSubsequence(str1, str2))
	})

	t.Run("test case 3", func(t *testing.T) {
		str1 := "ab"
		str2 := "d"

		assert.False(t, canMakeSubsequence(str1, str2))
	})
}
