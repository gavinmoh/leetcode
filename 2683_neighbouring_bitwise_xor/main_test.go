package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoesValidArrayExist(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		derived := []int{1, 1, 0}

		assert.True(t, doesValidArrayExist(derived))
	})

	t.Run("test case 2", func(t *testing.T) {
		derived := []int{1, 1}

		assert.True(t, doesValidArrayExist(derived))
	})

	t.Run("test case 3", func(t *testing.T) {
		derived := []int{1, 0}

		assert.False(t, doesValidArrayExist(derived))
	})
}
