package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseParentheses(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "(abcd)"
		expected := "dcba"

		assert.Equal(t, expected, reverseParentheses(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "(u(love)i)"
		expected := "iloveu"

		assert.Equal(t, expected, reverseParentheses(s))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "(ed(et(oc))el)"
		expected := "leetcode"

		assert.Equal(t, expected, reverseParentheses(s))
	})
}
