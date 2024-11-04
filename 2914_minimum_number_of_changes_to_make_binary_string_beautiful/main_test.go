package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinChanges(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "1001"
		expected := 2

		assert.Equal(t, expected, minChanges(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "10"
		expected := 1

		assert.Equal(t, expected, minChanges(s))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "0000"
		expected := 0

		assert.Equal(t, expected, minChanges(s))
	})
}
