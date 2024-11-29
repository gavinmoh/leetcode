package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIfExist(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		arr := []int{10, 2, 5, 3}

		assert.True(t, checkIfExist(arr))
	})

	t.Run("test case 2", func(t *testing.T) {
		arr := []int{3, 1, 7, 11}

		assert.False(t, checkIfExist(arr))
	})

	t.Run("test case 3", func(t *testing.T) {
		arr := []int{-2, 0, 10, -19, 4, 6, -8}

		assert.False(t, checkIfExist(arr))
	})

	t.Run("test case 4", func(t *testing.T) {
		arr := []int{0, 0}

		assert.True(t, checkIfExist(arr))
	})
}
