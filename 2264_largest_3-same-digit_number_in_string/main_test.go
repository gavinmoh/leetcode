package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestGoodInteger(t *testing.T) {
	t.Run("should return 777", func(t *testing.T) {
		num := "6777133339"
		expected := "777"
		assert.Equal(t, expected, largestGoodInteger(num))
	})

	t.Run("should return 000", func(t *testing.T) {
		num := "2300019"
		expected := "000"
		assert.Equal(t, expected, largestGoodInteger(num))
	})

	t.Run("should return empty string", func(t *testing.T) {
		num := "42352338"
		expected := ""
		assert.Equal(t, expected, largestGoodInteger(num))
	})
}
