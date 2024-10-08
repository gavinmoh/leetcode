package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSwaps(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "][]["
		expected := 1

		assert.Equal(t, expected, minSwaps(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "]]][[["
		expected := 2

		assert.Equal(t, expected, minSwaps(s))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "[]"
		expected := 0

		assert.Equal(t, expected, minSwaps(s))
	})
}
