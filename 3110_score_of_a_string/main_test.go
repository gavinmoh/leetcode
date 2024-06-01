package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreOfString(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "hello"
		expected := 13

		assert.Equal(t, expected, scoreOfString(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "zaz"
		expected := 50

		assert.Equal(t, expected, scoreOfString(s))
	})
}
