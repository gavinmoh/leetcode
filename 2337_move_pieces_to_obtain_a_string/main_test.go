package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanChange(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		start := "_L__R__R_"
		target := "L______RR"

		assert.True(t, canChange(start, target))
	})

	t.Run("test case 2", func(t *testing.T) {
		start := "R_L_"
		target := "__LR"

		assert.False(t, canChange(start, target))
	})

	t.Run("test case 3", func(t *testing.T) {
		start := "_R"
		target := "R_"

		assert.False(t, canChange(start, target))
	})
}
