package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfMatches(t *testing.T) {
	t.Run("should return 6", func(t *testing.T) {
		n := 7
		expected := 6
		assert.Equal(t, expected, numberOfMatches(n))
	})

	t.Run("should return 13", func(t *testing.T) {
		n := 14
		expected := 13
		assert.Equal(t, expected, numberOfMatches(n))
	})
}
