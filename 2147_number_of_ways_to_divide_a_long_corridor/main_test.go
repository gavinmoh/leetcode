package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfWays(t *testing.T) {
	t.Run("should return 3", func(t *testing.T) {
		corridor := "SSPPSPS"
		expected := 3
		assert.Equal(t, expected, numberOfWays(corridor))
	})

	t.Run("should return 1", func(t *testing.T) {
		corridor := "PPSPSP"
		expected := 1
		assert.Equal(t, expected, numberOfWays(corridor))
	})

	t.Run("should return 0", func(t *testing.T) {
		corridor := "S"
		expected := 0
		assert.Equal(t, expected, numberOfWays(corridor))
	})
}
