package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSteps(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 3
		expected := 3

		assert.Equal(t, expected, minSteps(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 1
		expected := 0

		assert.Equal(t, expected, minSteps(n))
	})
}
