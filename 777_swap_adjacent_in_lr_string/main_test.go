package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanTransform(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		start := "RXXLRXRXL"
		result := "XRLXXRRLX"

		assert.True(t, canTransform(start, result))
	})

	t.Run("test case 2", func(t *testing.T) {
		start := "X"
		result := "L"

		assert.False(t, canTransform(start, result))
	})
}
