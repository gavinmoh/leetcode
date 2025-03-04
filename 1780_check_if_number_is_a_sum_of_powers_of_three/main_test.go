package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPowersOfThree(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 12

		assert.True(t, checkPowersOfThree(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 91

		assert.True(t, checkPowersOfThree(n))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 21

		assert.False(t, checkPowersOfThree(n))
	})
}
