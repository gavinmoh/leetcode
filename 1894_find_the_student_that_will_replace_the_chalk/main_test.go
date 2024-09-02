package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChalkReplacer(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		chalk := []int{5, 1, 5}
		k := 22
		expected := 0

		assert.Equal(t, expected, chalkReplacer(chalk, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		chalk := []int{3, 4, 1, 2}
		k := 25
		expected := 1

		assert.Equal(t, expected, chalkReplacer(chalk, k))
	})
}
