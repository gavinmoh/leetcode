package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMaxDifference(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		num := 11891
		expected := 99009

		assert.Equal(t, expected, minMaxDifference(num))
	})

	t.Run("test case 2", func(t *testing.T) {
		num := 90
		expected := 99

		assert.Equal(t, expected, minMaxDifference(num))
	})
}
