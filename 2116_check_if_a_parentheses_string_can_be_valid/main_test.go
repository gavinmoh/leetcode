package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanBeValid(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "))()))"
		locked := "010100"

		assert.True(t, canBeValid(s, locked))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "()()"
		locked := "0000"

		assert.True(t, canBeValid(s, locked))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := ")"
		locked := "0"

		assert.False(t, canBeValid(s, locked))
	})

	t.Run("test case 4", func(t *testing.T) {
		s := ")("
		locked := "00"

		assert.True(t, canBeValid(s, locked))
	})

	t.Run("test case 5", func(t *testing.T) {
		s := "()))"
		locked := "0010"

		assert.True(t, canBeValid(s, locked))
	})

	t.Run("test case 6", func(t *testing.T) {
		s := "())()))()(()(((())(()()))))((((()())(())"
		locked := "1011101100010001001011000000110010100101"

		assert.True(t, canBeValid(s, locked))
	})

	t.Run("test case 7", func(t *testing.T) {
		s := "())(()(()(())()())(())((())(()())((())))))(((((((())(()))))("
		locked := "100011110110011011010111100111011101111110000101001101001111"

		assert.False(t, canBeValid(s, locked))
	})
}
