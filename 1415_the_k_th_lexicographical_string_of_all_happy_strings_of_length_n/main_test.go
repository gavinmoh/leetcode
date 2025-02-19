package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHappyString(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 1
		k := 3
		expected := "c"

		assert.Equal(t, expected, getHappyString(n, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 1
		k := 4
		expected := ""

		assert.Equal(t, expected, getHappyString(n, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 3
		k := 9
		expected := "cab"

		assert.Equal(t, expected, getHappyString(n, k))
	})
}
