package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonChars(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		words := []string{"bella", "label", "roller"}
		expected := []string{"l", "l", "e"}

		assert.Equal(t, expected, commonChars(words))
	})

	t.Run("test case 2", func(t *testing.T) {
		words := []string{"cool", "lock", "cook"}
		expected := []string{"c", "o"}

		assert.Equal(t, expected, commonChars(words))
	})
}
