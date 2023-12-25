package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumDecodings(t *testing.T) {
	t.Run("should return 2", func(t *testing.T) {
		s := "12"
		expected := 2
		assert.Equal(t, expected, numDecodings(s))
	})

	t.Run("should return 3", func(t *testing.T) {
		s := "226"
		expected := 3
		assert.Equal(t, expected, numDecodings(s))
	})

	t.Run("should return 0", func(t *testing.T) {
		s := "06"
		expected := 0
		assert.Equal(t, expected, numDecodings(s))
	})

	t.Run("should return 1", func(t *testing.T) {
		s := "27"
		expected := 1
		assert.Equal(t, expected, numDecodings(s))
	})

}
