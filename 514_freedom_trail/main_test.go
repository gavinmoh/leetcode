package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRotateSteps(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		ring := "godding"
		key := "gd"
		expected := 4

		assert.Equal(t, expected, findRotateSteps(ring, key))
	})

	t.Run("test case 2", func(t *testing.T) {
		ring := "godding"
		key := "godding"
		expected := 13

		assert.Equal(t, expected, findRotateSteps(ring, key))
	})
}
