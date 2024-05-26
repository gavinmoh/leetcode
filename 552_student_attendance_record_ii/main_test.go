package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckRecord(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 2
		expected := 8

		assert.Equal(t, expected, checkRecord(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 1
		expected := 3

		assert.Equal(t, expected, checkRecord(n))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 10101
		expected := 183236316

		assert.Equal(t, expected, checkRecord(n))
	})

}
