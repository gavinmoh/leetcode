package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBinary(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		a := "11"
		b := "1"
		expected := "100"

		assert.Equal(t, expected, addBinary(a, b))
	})

	t.Run("test case 2", func(t *testing.T) {
		a := "1010"
		b := "1011"
		expected := "10101"

		assert.Equal(t, expected, addBinary(a, b))
	})

	t.Run("test case 3", func(t *testing.T) {
		a := "1111"
		b := "1101000100"
		expected := "1101010011"

		assert.Equal(t, expected, addBinary(a, b))
	})

	t.Run("test case 4", func(t *testing.T) {
		a := "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
		b := "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
		expected := "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000"

		assert.Equal(t, expected, addBinary(a, b))
	})
}
