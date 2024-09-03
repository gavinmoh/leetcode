package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLucky(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "iiii"
		k := 1
		expected := 36

		assert.Equal(t, expected, getLucky(s, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "leetcode"
		k := 2
		expected := 6

		assert.Equal(t, expected, getLucky(s, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "zbax"
		k := 2
		expected := 8

		assert.Equal(t, expected, getLucky(s, k))
	})
}
