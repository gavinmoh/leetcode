package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBoolExpr(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		expression := "&(|(f))"

		assert.False(t, parseBoolExpr(expression))
	})

	t.Run("test case 2", func(t *testing.T) {
		expression := "|(f,f,f,t)"

		assert.True(t, parseBoolExpr(expression))
	})

	t.Run("test case 3", func(t *testing.T) {
		expression := "!(&(f,t))"

		assert.True(t, parseBoolExpr(expression))
	})
}
