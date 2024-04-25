package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestIdealString(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "acfgbd"
		k := 2
		expected := 4

		assert.Equal(t, expected, longestIdealString(s, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "abcd"
		k := 3
		expected := 4

		assert.Equal(t, expected, longestIdealString(s, k))
	})
}
