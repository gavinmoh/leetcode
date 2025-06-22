package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivideString(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "abcdefghi"
		k := 3
		fill := byte('x')
		expected := []string{"abc", "def", "ghi"}

		assert.Equal(t, expected, divideString(s, k, fill))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "abcdefghij"
		k := 3
		fill := byte('x')
		expected := []string{"abc", "def", "ghi", "jxx"}

		assert.Equal(t, expected, divideString(s, k, fill))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "ctoyjrwtngqwt"
		k := 8
		fill := byte('n')
		expected := []string{"ctoyjrwt", "ngqwtnnn"}

		assert.Equal(t, expected, divideString(s, k, fill))
	})
}
