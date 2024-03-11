package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomSortString(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		order := "cba"
		s := "abcd"
		expected := "cbad"

		assert.Equal(t, expected, customSortString(order, s))
	})

	t.Run("test case 2", func(t *testing.T) {
		order := "bcafg"
		s := "abcd"
		expected := "bcad"

		assert.Equal(t, expected, customSortString(order, s))
	})
}
