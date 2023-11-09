package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountHomogenous(t *testing.T) {
	t.Run("should return 13", func(t *testing.T) {
		s := "abbcccaa"
		expected := 13
		assert.Equal(t, expected, countHomogenous(s))
	})

	t.Run("should return 2", func(t *testing.T) {
		s := "xy"
		expected := 2
		assert.Equal(t, expected, countHomogenous(s))
	})

	t.Run("should return 15", func(t *testing.T) {
		s := "zzzzz"
		expected := 15
		assert.Equal(t, expected, countHomogenous(s))
	})
}
