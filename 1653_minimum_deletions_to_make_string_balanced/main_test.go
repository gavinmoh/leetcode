package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumDeletions(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "aababbab"
		expected := 2

		assert.Equal(t, expected, minimumDeletions(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "bbaaaaabb"
		expected := 2

		assert.Equal(t, expected, minimumDeletions(s))
	})
}
