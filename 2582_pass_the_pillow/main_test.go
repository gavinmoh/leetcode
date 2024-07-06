package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassThePillow(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 4
		time := 5
		expected := 2

		assert.Equal(t, expected, passThePillow(n, time))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 3
		time := 2
		expected := 3

		assert.Equal(t, expected, passThePillow(n, time))
	})
}
