package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanBeEqual(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		target := []int{1, 2, 3, 4}
		arr := []int{2, 4, 1, 3}

		assert.True(t, canBeEqual(target, arr))
	})

	t.Run("test case 2", func(t *testing.T) {
		target := []int{7}
		arr := []int{7}

		assert.True(t, canBeEqual(target, arr))
	})

	t.Run("test case 3", func(t *testing.T) {
		target := []int{3, 7, 9}
		arr := []int{3, 7, 11}

		assert.False(t, canBeEqual(target, arr))
	})
}
