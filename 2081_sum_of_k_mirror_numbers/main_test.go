package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKMirror(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		k := 2
		n := 5
		expected := int64(25)

		assert.Equal(t, expected, kMirror(k, n))
	})

	t.Run("test case 2", func(t *testing.T) {
		k := 3
		n := 7
		expected := int64(499)

		assert.Equal(t, expected, kMirror(k, n))
	})

	t.Run("test case 3", func(t *testing.T) {
		k := 7
		n := 17
		expected := int64(20379000)

		assert.Equal(t, expected, kMirror(k, n))
	})
}
