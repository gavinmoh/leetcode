package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeEqual(t *testing.T) {
	t.Run("should return true", func(t *testing.T) {
		words := []string{"abc", "aabc", "bc"}
		assert.True(t, makeEqual(words))
	})

	t.Run("should return false", func(t *testing.T) {
		words := []string{"ab", "a"}
		assert.False(t, makeEqual(words))
	})
}
