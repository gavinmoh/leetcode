package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinRemoveToMakeValid(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "lee(t(c)o)de)"
		expected := "lee(t(c)o)de"

		assert.Equal(t, expected, minRemoveToMakeValid(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "a)b(c)d"
		expected := "ab(c)d"

		assert.Equal(t, expected, minRemoveToMakeValid(s))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "))(("
		expected := ""

		assert.Equal(t, expected, minRemoveToMakeValid(s))
	})

	t.Run("test case 4", func(t *testing.T) {
		s := "(a(b(c)d)"
		expected := "(a(bc)d)"

		assert.Equal(t, expected, minRemoveToMakeValid(s))
	})

	t.Run("test case 5", func(t *testing.T) {
		s := "())()((("
		expected := "()()"

		assert.Equal(t, expected, minRemoveToMakeValid(s))
	})
}
