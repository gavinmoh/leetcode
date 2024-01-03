package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfBeams(t *testing.T) {
	t.Run("should return 8", func(t *testing.T) {
		bank := []string{"011001", "000000", "010100", "001000"}
		expected := 8
		assert.Equal(t, expected, numberOfBeams(bank))
	})

	t.Run("should return 0", func(t *testing.T) {
		bank := []string{"000", "111", "000"}
		expected := 0
		assert.Equal(t, expected, numberOfBeams(bank))
	})
}
