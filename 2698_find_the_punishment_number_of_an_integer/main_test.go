package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPunishmentNumber(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 10
		expected := 182

		assert.Equal(t, expected, punishmentNumber(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 37
		expected := 1478

		assert.Equal(t, expected, punishmentNumber(n))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 45
		expected := 3503

		assert.Equal(t, expected, punishmentNumber(n))
	})

	t.Run("test case 4", func(t *testing.T) {
		n := 235
		expected := 96559

		assert.Equal(t, expected, punishmentNumber(n))
	})
}
