package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTheLongestSubstring(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "eleetminicoworoep"
		expected := 13

		assert.Equal(t, expected, findTheLongestSubstring(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "leetcodeisgreat"
		expected := 5

		assert.Equal(t, expected, findTheLongestSubstring(s))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "bcbcbc"
		expected := 6

		assert.Equal(t, expected, findTheLongestSubstring(s))
	})
}
