package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumWaterBottles(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		numBottles := 9
		numExchange := 3
		expected := 13

		assert.Equal(t, expected, numWaterBottles(numBottles, numExchange))
	})

	t.Run("test case 2", func(t *testing.T) {
		numBottles := 15
		numExchange := 4
		expected := 19

		assert.Equal(t, expected, numWaterBottles(numBottles, numExchange))
	})
}
