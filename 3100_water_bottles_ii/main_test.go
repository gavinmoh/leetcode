package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxBottlesDrunk(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		numBottles := 13
		numExchange := 6
		expected := 15

		assert.Equal(t, expected, maxBottlesDrunk(numBottles, numExchange))
	})

	t.Run("test case 2", func(t *testing.T) {
		numBottles := 10
		numExchange := 3
		expected := 13

		assert.Equal(t, expected, maxBottlesDrunk(numBottles, numExchange))
	})

	t.Run("test case 3", func(t *testing.T) {
		numBottles := 13
		numExchange := 13
		expected := 14

		assert.Equal(t, expected, maxBottlesDrunk(numBottles, numExchange))
	})
}
