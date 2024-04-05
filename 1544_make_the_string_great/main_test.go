package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeGood(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "leEeetcode"
		expected := "leetcode"

		assert.Equal(t, expected, makeGood(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "abBAcC"
		expected := ""

		assert.Equal(t, expected, makeGood(s))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "s"
		expected := "s"

		assert.Equal(t, expected, makeGood(s))
	})
}
