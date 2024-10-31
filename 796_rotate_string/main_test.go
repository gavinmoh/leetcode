package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateString(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "abcde"
		goal := "cdeab"

		assert.True(t, rotateString(s, goal))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "abcde"
		goal := "abced"

		assert.False(t, rotateString(s, goal))
	})
}
