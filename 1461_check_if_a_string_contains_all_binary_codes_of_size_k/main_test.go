package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasAllCodes(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "00110110"
		k := 2

		assert.True(t, hasAllCodes(s, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "0110"
		k := 1

		assert.True(t, hasAllCodes(s, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "0110"
		k := 2

		assert.False(t, hasAllCodes(s, k))
	})

	t.Run("test case 4", func(t *testing.T) {
		s := "00110"
		k := 2

		assert.True(t, hasAllCodes(s, k))
	})
}
