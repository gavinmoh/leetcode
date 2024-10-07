package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinLength(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "ABFCACDB"
		expected := 2

		assert.Equal(t, expected, minLength(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "ACBBD"
		expected := 5

		assert.Equal(t, expected, minLength(s))
	})
}
