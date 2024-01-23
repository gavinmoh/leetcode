package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxLength(t *testing.T) {
	t.Run("should return 4", func(t *testing.T) {
		arr := []string{"un", "iq", "ue"}
		expected := 4
		assert.Equal(t, expected, maxLength(arr))
	})

	t.Run("should return 6", func(t *testing.T) {
		arr := []string{"cha", "r", "act", "ers"}
		expected := 6
		assert.Equal(t, expected, maxLength(arr))
	})

	t.Run("should return 26", func(t *testing.T) {
		arr := []string{"abcdefghijklmnopqrstuvwxyz"}
		expected := 26
		assert.Equal(t, expected, maxLength(arr))
	})

	t.Run("should return 0", func(t *testing.T) {
		arr := []string{"aa", "bb"}
		expected := 0
		assert.Equal(t, expected, maxLength(arr))
	})
}
