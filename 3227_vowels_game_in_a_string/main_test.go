package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoesAliceWin(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "leetcoder"

		assert.True(t, doesAliceWin(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "bbcd"

		assert.False(t, doesAliceWin(s))
	})
}
