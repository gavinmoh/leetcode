package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumPushes(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		word := "abcde"
		expected := 5

		assert.Equal(t, expected, minimumPushes(word))
	})

	t.Run("test case 2", func(t *testing.T) {
		word := "xyzxyzxyzxyz"
		expected := 12

		assert.Equal(t, expected, minimumPushes(word))
	})

	t.Run("test case 3", func(t *testing.T) {
		word := "aabbccddeeffgghhiiiiii"
		expected := 24

		assert.Equal(t, expected, minimumPushes(word))
	})
}
