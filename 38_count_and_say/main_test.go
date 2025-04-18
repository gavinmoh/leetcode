package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAndSay(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 4
		expected := "1211"

		assert.Equal(t, expected, countAndSay(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 1
		expected := "1"

		assert.Equal(t, expected, countAndSay(n))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 5
		expected := "111221"

		assert.Equal(t, expected, countAndSay(n))
	})
}
