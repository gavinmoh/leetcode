package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindComplement(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		num := 5
		expected := 2

		assert.Equal(t, expected, findComplement(num))
	})

	t.Run("test case 2", func(t *testing.T) {
		num := 1
		expected := 0

		assert.Equal(t, expected, findComplement(num))
	})

	t.Run("test case 3", func(t *testing.T) {
		num := 2
		expected := 1

		assert.Equal(t, expected, findComplement(num))
	})
}
