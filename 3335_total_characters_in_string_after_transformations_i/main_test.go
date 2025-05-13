package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthAfterTransformations(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "abcyy"
		transform := 2
		expected := 7

		assert.Equal(t, expected, lengthAfterTransformations(s, transform))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "azbk"
		transform := 1
		expected := 5

		assert.Equal(t, expected, lengthAfterTransformations(s, transform))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "jqktcurgdvlibczdsvnsg"
		transform := 7517
		expected := 79033769

		assert.Equal(t, expected, lengthAfterTransformations(s, transform))
	})
}
