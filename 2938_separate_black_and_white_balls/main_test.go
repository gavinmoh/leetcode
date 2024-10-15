package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumSteps(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "101"
		expected := int64(1)

		assert.Equal(t, expected, minimumSteps(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "100"
		expected := int64(2)

		assert.Equal(t, expected, minimumSteps(s))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "0111"
		expected := int64(0)

		assert.Equal(t, expected, minimumSteps(s))
	})
}
