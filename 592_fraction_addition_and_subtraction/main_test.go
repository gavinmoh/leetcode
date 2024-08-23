package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFractionAddition(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		expression := "-1/2+1/2"
		expected := "0/1"

		assert.Equal(t, expected, fractionAddition(expression))
	})

	t.Run("test case 2", func(t *testing.T) {
		expression := "-1/2+1/2+1/3"
		expected := "1/3"

		assert.Equal(t, expected, fractionAddition(expression))
	})

	t.Run("test case 3", func(t *testing.T) {
		expression := "1/3-1/2"
		expected := "-1/6"

		assert.Equal(t, expected, fractionAddition(expression))
	})
}
