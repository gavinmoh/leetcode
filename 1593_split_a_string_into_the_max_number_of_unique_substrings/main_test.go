package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxUniqueSplit(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "ababccc"
		expected := 5

		assert.Equal(t, expected, maxUniqueSplit(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "aba"
		expected := 2

		assert.Equal(t, expected, maxUniqueSplit(s))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "aa"
		expected := 1

		assert.Equal(t, expected, maxUniqueSplit(s))
	})
}
