package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountLargestGroup(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 13
		expected := 4

		assert.Equal(t, expected, countLargestGroup(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 2
		expected := 2

		assert.Equal(t, expected, countLargestGroup(n))
	})
}
