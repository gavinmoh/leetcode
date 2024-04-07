package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckValidString(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "()"

		assert.True(t, checkValidString(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "(*)"

		assert.True(t, checkValidString(s))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "(*))"

		assert.True(t, checkValidString(s))
	})

	t.Run("test case 4", func(t *testing.T) {
		s := "("

		assert.False(t, checkValidString(s))
	})

	t.Run("test case 5", func(t *testing.T) {
		s := "(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())"

		assert.False(t, checkValidString(s))
	})
}
