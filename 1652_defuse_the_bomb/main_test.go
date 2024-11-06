package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecrypt(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		code := []int{5, 7, 1, 4}
		k := 3
		expected := []int{12, 10, 16, 13}

		assert.Equal(t, expected, decrypt(code, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		code := []int{1, 2, 3, 4}
		k := 0
		expected := []int{0, 0, 0, 0}

		assert.Equal(t, expected, decrypt(code, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		code := []int{2, 4, 9, 3}
		k := -2
		expected := []int{12, 5, 6, 13}

		assert.Equal(t, expected, decrypt(code, k))
	})
}
